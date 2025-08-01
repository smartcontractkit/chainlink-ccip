use anchor_lang::prelude::*;

use crate::{messages::ExecutionReportSingleChain, CcipOfframpError, ExecutionReportBuffer};
pub trait Buffering {
    fn is_initialized(&self) -> bool;
    fn filled_chunks(&self) -> u8;
    fn is_complete(&self) -> bool;
    fn bytes(&self) -> Result<&[u8]>;
    fn add_chunk(
        &mut self,
        report_length: u32,
        chunk: &[u8],
        chunk_index: u8,
        num_chunks: u8,
    ) -> Result<()>;
}

impl Buffering for ExecutionReportBuffer {
    fn is_initialized(&self) -> bool {
        !self.data.is_empty()
    }

    fn filled_chunks(&self) -> u8 {
        self.chunk_bitmap.count_ones() as u8
    }

    fn is_complete(&self) -> bool {
        self.is_initialized() && self.filled_chunks() == self.num_chunks
    }

    fn bytes(&self) -> Result<&[u8]> {
        require!(
            self.is_complete(),
            CcipOfframpError::ExecutionReportBufferIncomplete
        );

        Ok(&self.data)
    }

    fn add_chunk(
        &mut self,
        report_length: u32,
        chunk: &[u8],
        chunk_index: u8,
        num_chunks: u8,
    ) -> Result<()> {
        require!(
            num_chunks > 0 && chunk_index < num_chunks,
            CcipOfframpError::ExecutionReportBufferInvalidChunkIndex
        );
        require!(
            !chunk.is_empty() && chunk.len() <= report_length as usize,
            CcipOfframpError::ExecutionReportBufferInvalidChunkSize
        );

        if !self.is_initialized() {
            self.initialize(report_length, chunk.len() as u32, num_chunks, chunk_index)?;
        }

        require_eq!(
            self.data.len(),
            report_length as usize,
            CcipOfframpError::ExecutionReportBufferInvalidLength
        );
        require_eq!(
            self.num_chunks,
            num_chunks,
            CcipOfframpError::ExecutionReportBufferInvalidChunkNumber
        );

        let chunk_mask = 1u64 << chunk_index;
        require!(
            chunk_mask & self.chunk_bitmap == 0,
            CcipOfframpError::ExecutionReportBufferAlreadyContainsChunk
        );

        require_gte!(
            self.chunk_length,
            chunk.len() as u32,
            CcipOfframpError::ExecutionReportBufferInvalidChunkSize
        );

        if chunk.len() < self.chunk_length as usize {
            // Only the terminator (last chunk) can be smaller than the others.
            require_eq!(
                chunk_index,
                self.num_chunks - 1,
                CcipOfframpError::ExecutionReportBufferInvalidChunkSize
            );
        }

        require_gt!(
            self.num_chunks,
            chunk_index,
            CcipOfframpError::ExecutionReportBufferInvalidChunkIndex
        );

        let start = self.chunk_length as usize * chunk_index as usize;
        let end = self.data.len().min(start + chunk.len());
        self.data[start..end].copy_from_slice(chunk);
        self.chunk_bitmap |= chunk_mask;

        Ok(())
    }
}

impl ExecutionReportBuffer {
    fn initialize(
        &mut self,
        report_length: u32,
        chunk_length: u32,
        total_chunks: u8,
        chunk_index: u8,
    ) -> Result<()> {
        require!(
            !self.is_initialized(),
            CcipOfframpError::ExecutionReportBufferAlreadyInitialized
        );
        require!(
            report_length > 0 && chunk_length <= report_length && chunk_length > 0,
            CcipOfframpError::ExecutionReportBufferInvalidLength
        );
        require!(
            total_chunks <= 64 && total_chunks > 0,
            CcipOfframpError::ExecutionReportBufferInvalidChunkSize
        );

        // If we're initializing with the last chunk, it could be smaller
        // than the rest, so we calculate the chunk size based on the expected
        // size of all the others.
        let is_last = chunk_index == total_chunks - 1;
        let global_chunk_length = if is_last && total_chunks > 1 {
            (report_length - chunk_length) / (total_chunks as u32 - 1)
        } else {
            chunk_length
        };

        require_gt!(
            global_chunk_length,
            0,
            CcipOfframpError::ExecutionReportBufferInvalidChunkSize
        );

        require_eq!(
            total_chunks as u32,
            div_ceil(report_length, global_chunk_length),
            CcipOfframpError::ExecutionReportBufferInvalidLength,
        );
        self.chunk_length = global_chunk_length;
        self.num_chunks = total_chunks;
        self.data.resize(report_length as usize, 0);

        Ok(())
    }
}

pub fn deserialize_from_buffer_account(
    execution_report_buffer: &AccountInfo,
) -> Result<(ExecutionReportSingleChain, usize)> {
    // Ensures the buffer is initialized, and owned by the program.
    require_keys_eq!(
        *execution_report_buffer.owner,
        crate::ID,
        CcipOfframpError::ExecutionReportUnavailable
    );
    let buffer = ExecutionReportBuffer::try_deserialize(
        &mut execution_report_buffer.data.borrow().as_ref(),
    )?;

    Ok((
        ExecutionReportSingleChain::deserialize(&mut buffer.bytes()?)
            .map_err(|_| CcipOfframpError::FailedToDeserializeReport)?,
        buffer.data.len(),
    ))
}

// Using a.div_ceil(b) is unstable on Rust 1.68, so we implement our own version even though clippy may complain.
// See https://github.com/rust-lang/rust/issues/88581 for more info.
// Whenever we upgrade Anchor & Rust, we can remove this.
fn div_ceil<T: Into<u32>>(a: T, b: T) -> u32 {
    let (a, b) = (a.into(), b.into());
    assert!(a + b - 1 > 0);

    // on newer rust version, this would require #[allow(clippy::manual_div_ceil)]
    // or be replaced with `a.div_ceil(b)` directly. However, on older rust versions,
    // there is no support for clippy's lint and the div_ceil method is unstable.
    (a + b - 1) / b
}

#[cfg(test)]
mod tests {
    use super::*;

    fn empty_buffer() -> ExecutionReportBuffer {
        ExecutionReportBuffer {
            data: vec![],
            chunk_bitmap: 0,
            num_chunks: 0,
            chunk_length: 0,
            version: 0,
        }
    }

    #[test]
    fn rejects_invalid_initializations() {
        let mut buffer = empty_buffer();

        assert_eq!(
            buffer.initialize(10, 100, 1, 0).unwrap_err(),
            CcipOfframpError::ExecutionReportBufferInvalidLength.into()
        );

        let mut buffer = empty_buffer();
        assert_eq!(
            buffer.initialize(0, 1, 1, 0).unwrap_err(),
            CcipOfframpError::ExecutionReportBufferInvalidLength.into()
        );

        let mut buffer = empty_buffer();
        // Too small buffer sizes aren't OK as they'd break the bitmap
        assert_eq!(
            buffer.initialize(100000, 1, 1, 0).unwrap_err(),
            CcipOfframpError::ExecutionReportBufferInvalidLength.into()
        );

        let mut buffer = empty_buffer();
        assert_eq!(
            buffer.initialize(100, 10, 255, 0).unwrap_err(),
            CcipOfframpError::ExecutionReportBufferInvalidChunkSize.into()
        );
    }

    #[test]
    fn buffering_happy_path() {
        let mut buffer = empty_buffer();

        const DIVISIBLE_BY_THREE_SIZE: usize = 33;

        let message: &[u8; DIVISIBLE_BY_THREE_SIZE] = b"Very large message, wow so large.";
        let chunk_size = DIVISIBLE_BY_THREE_SIZE / 3;

        for i in 0..3 {
            assert!(!buffer.is_complete());
            buffer
                .add_chunk(
                    message.len() as u32,
                    &message[i * chunk_size..(i + 1) * chunk_size],
                    i as u8,
                    3,
                )
                .unwrap();
            assert_eq!(buffer.filled_chunks() as usize, i + 1);
        }

        dbg!(&buffer);
        assert!(buffer.is_complete());
        assert_eq!(buffer.bytes().unwrap(), message);
    }

    #[test]
    fn buffering_with_a_smaller_terminator() {
        let mut buffer = empty_buffer();

        const NOT_DIVISIBLE_BY_THREE_SIZE: usize = 35;

        let message: &[u8; NOT_DIVISIBLE_BY_THREE_SIZE] = b"Very large message, wow so large!!!";
        let chunk_size = NOT_DIVISIBLE_BY_THREE_SIZE / 3;

        for i in 0..4 {
            assert!(!buffer.is_complete());
            buffer
                .add_chunk(
                    message.len() as u32,
                    &message[i * chunk_size..message.len().min((i + 1) * chunk_size)],
                    i as u8,
                    4,
                )
                .unwrap();
            assert_eq!(buffer.filled_chunks() as usize, i + 1);
        }

        dbg!(std::str::from_utf8(&buffer.data).unwrap());
        assert!(buffer.is_complete());
        assert_eq!(buffer.bytes().unwrap(), message);
    }

    #[test]
    fn buffering_with_a_smaller_terminator_out_of_order() {
        let mut buffer = empty_buffer();

        const NOT_DIVISIBLE_BY_THREE_SIZE: usize = 35;

        let message: &[u8; NOT_DIVISIBLE_BY_THREE_SIZE] = b"Very large message, wow so large!!!";
        let chunk_size = NOT_DIVISIBLE_BY_THREE_SIZE / 3;

        // Note the rev(). The smaller terminator will be received first!
        for i in (0..4).rev() {
            assert!(!buffer.is_complete());
            buffer
                .add_chunk(
                    message.len() as u32,
                    &message[i * chunk_size..message.len().min((i + 1) * chunk_size)],
                    i as u8,
                    4,
                )
                .unwrap();
            assert_eq!(buffer.filled_chunks() as usize, 4 - i);
        }

        dbg!(&buffer);
        dbg!("{}", std::str::from_utf8(&buffer.data).unwrap());
        assert!(buffer.is_complete());
        assert_eq!(buffer.bytes().unwrap(), message);
    }

    #[test]
    fn rejects_invalid_chunks() {
        let mut buffer = empty_buffer();
        buffer.add_chunk(22, b"AAAAA", 0, 5).unwrap();
        // This is OK as it is the terminator.
        buffer.add_chunk(22, b"BB", 4, 5).unwrap();
        // This is not OK.
        buffer.add_chunk(22, b"CCCCCCCCCCC", 2, 5).unwrap_err();

        let mut buffer = empty_buffer();
        buffer.add_chunk(100, b"small", 0, 20).unwrap();
        // Invalid due to inconsistent total report size.
        buffer.add_chunk(50, b"small", 1, 20).unwrap_err();

        let mut buffer = empty_buffer();
        buffer.add_chunk(100, b"small", 0, 20).unwrap();
        // Invalid due to repeated index.
        buffer.add_chunk(100, b"small", 0, 20).unwrap_err();

        let mut buffer = empty_buffer();
        buffer
            .add_chunk(10, b"Much bigger than ten characters", 0, 1)
            .unwrap_err();

        let mut buffer = empty_buffer();
        buffer.add_chunk(29, b"Medium sized", 1, 3).unwrap();
        // Cannot be smaller: only the terminator can (must come at the end.)
        buffer.add_chunk(29, b"small", 0, 3).unwrap_err();
        buffer.add_chunk(29, b"small", 2, 3).unwrap();
    }
}
