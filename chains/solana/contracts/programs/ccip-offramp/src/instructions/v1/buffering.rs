use anchor_lang::prelude::*;
use ccip_common::seed;

use crate::{messages::ExecutionReportSingleChain, CcipOfframpError, ExecutionReportBuffer};
pub trait Buffering {
    fn is_initialized(&self) -> bool;
    fn filled_chunks(&self) -> u32;
    fn is_complete(&self) -> bool;
    fn bytes(&self) -> Result<&[u8]>;
    fn initialize(&mut self, report_length: u32, chunk_length: u32) -> Result<()>;
    fn add_chunk(&mut self, report_length: u32, chunk: &[u8], chunk_index: u8) -> Result<()>;
    fn recover_wrong_size(
        &mut self,
        report_length: u32,
        new_chunk: &[u8],
        chunk_index: u8,
    ) -> Result<()>;
}

impl Buffering for ExecutionReportBuffer {
    fn is_initialized(&self) -> bool {
        self.total_chunks > 0
    }

    fn filled_chunks(&self) -> u32 {
        self.chunk_bitmap.count_ones()
    }

    fn is_complete(&self) -> bool {
        self.is_initialized() && self.filled_chunks() == self.total_chunks
    }

    fn bytes(&self) -> Result<&[u8]> {
        require!(
            self.is_complete(),
            CcipOfframpError::ExecutionReportBufferIncomplete
        );

        Ok(&self.data)
    }

    fn initialize(&mut self, report_length: u32, chunk_length: u32) -> Result<()> {
        require!(
            !self.is_initialized(),
            CcipOfframpError::ExecutionReportBufferAlreadyInitialized
        );
        require!(
            chunk_length > 0 && report_length >= chunk_length,
            CcipOfframpError::ExecutionReportBufferInvalidLength
        );
        self.data.resize(report_length as usize, 0);
        self.total_chunks = (report_length + chunk_length - 1) / chunk_length;
        require_gt!(
            64,
            self.total_chunks,
            CcipOfframpError::ExecutionReportBufferChunkSizeTooSmall
        );
        self.chunk_length = chunk_length;
        self.report_length = report_length;
        Ok(())
    }

    fn add_chunk(&mut self, report_length: u32, chunk: &[u8], chunk_index: u8) -> Result<()> {
        if !self.is_initialized() {
            self.initialize(report_length, chunk.len() as u32)?;
        }

        require_eq!(
            self.report_length,
            report_length,
            CcipOfframpError::ExecutionReportBufferInvalidLength
        );

        let chunk_mask = 1u64 << chunk_index;
        require!(
            chunk_mask & self.chunk_bitmap == 0,
            CcipOfframpError::ExecutionReportBufferAlreadyContainsChunk
        );

        if chunk.len() as u32 > self.chunk_length {
            // We hit the special case where the first received chunk was the last one
            // in the buffer (terminator), which may be smaller than all others. It's okay,
            // we can recover in place.
            return self.recover_wrong_size(report_length, chunk, chunk_index);
        }

        require_gte!(
            self.chunk_length,
            chunk.len() as u32,
            CcipOfframpError::ExecutionReportBufferInvalidChunkSize
        );

        if chunk.len() < self.chunk_length as usize {
            // Only the terminator (last chunk) can be smaller than the others.
            require_eq!(
                chunk_index as u32,
                self.total_chunks - 1,
                CcipOfframpError::ExecutionReportBufferInvalidChunkSize
            );
        }

        require_gt!(
            self.total_chunks,
            chunk_index as u32,
            CcipOfframpError::ExecutionReportBufferInvalidChunkIndex
        );

        let start = self.chunk_length as usize * chunk_index as usize;
        let end = self.data.len().min(start + chunk.len());
        self.data[start..end].copy_from_slice(chunk);
        self.chunk_bitmap |= chunk_mask;

        Ok(())
    }

    fn recover_wrong_size(
        &mut self,
        report_length: u32,
        new_chunk: &[u8],
        chunk_index: u8,
    ) -> Result<()> {
        // Only makes sense to recover if we got the first chunk wrong (because it was the buffer
        // terminator). Any size mismatch beyond that means the user is sending the chunks incorrectly.
        require_eq!(
            self.filled_chunks(),
            1,
            CcipOfframpError::ExecutionReportBufferInvalidChunkSize
        );

        // We extract what we now know is the terminator
        let terminator_index = self.chunk_bitmap.trailing_zeros() as u8;
        let mut terminator = vec![0u8; self.chunk_length as usize];
        let start = terminator_index as usize * self.chunk_length as usize;
        let end = start + terminator.len();
        terminator.copy_from_slice(&self.data[start..end]);

        // We reset the buffer metadata. It's okay to leave the old data in, as it will be clobbered.
        self.chunk_bitmap = 0;
        self.total_chunks = 0;
        self.chunk_length = 0;

        // We reinsert the new chunk and terminator, which will be accepted as it's smaller. From now
        // on, we won't accept bigger chunks than this again.
        self.add_chunk(report_length, new_chunk, chunk_index)?;
        self.add_chunk(report_length, &terminator, terminator_index)?;
        Ok(())
    }
}

pub fn deserialize_from_buffer_account(
    execution_report_buffer: &AccountInfo,
    authority: Pubkey,
    merkle_root: &[u8],
) -> Result<(ExecutionReportSingleChain, usize)> {
    // Ensures the buffer is initialized, and owned by the program.
    require_keys_eq!(
        *execution_report_buffer.owner,
        crate::ID,
        CcipOfframpError::ExecutionReportUnavailable
    );
    let (expected_buffer_key, _) = Pubkey::find_program_address(
        &[
            seed::EXECUTION_REPORT_BUFFER,
            merkle_root,
            authority.as_ref(),
        ],
        &crate::ID,
    );
    require_keys_eq!(
        expected_buffer_key,
        execution_report_buffer.key(),
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

#[cfg(test)]
mod tests {
    use super::*;

    fn empty_buffer() -> ExecutionReportBuffer {
        ExecutionReportBuffer {
            data: vec![],
            chunk_bitmap: 0,
            total_chunks: 0,
            chunk_length: 0,
            report_length: 0,
            version: 0,
        }
    }

    #[test]
    fn rejects_invalid_initializations() {
        let mut buffer = empty_buffer();

        assert_eq!(
            buffer.initialize(10, 100).unwrap_err(),
            CcipOfframpError::ExecutionReportBufferInvalidLength.into()
        );

        assert_eq!(
            buffer.initialize(0, 1).unwrap_err(),
            CcipOfframpError::ExecutionReportBufferInvalidLength.into()
        );

        // Too small buffer sizes aren't OK as they'd break the bitmap
        assert_eq!(
            buffer.initialize(100000, 1).unwrap_err(),
            CcipOfframpError::ExecutionReportBufferChunkSizeTooSmall.into()
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
                )
                .unwrap();
            assert_eq!(buffer.filled_chunks() as usize, i + 1);
        }

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
        buffer.add_chunk(22, b"AAAAA", 0).unwrap();
        // This is OK as it is the terminator.
        buffer.add_chunk(22, b"BB", 4).unwrap();
        // This is not OK.
        buffer.add_chunk(22, b"CCCCCCCCCCC", 2).unwrap_err();

        let mut buffer = empty_buffer();
        buffer.add_chunk(100, b"small", 0).unwrap();
        // Invalid due to inconsistent total report size.
        buffer.add_chunk(50, b"small", 1).unwrap_err();

        let mut buffer = empty_buffer();
        buffer.add_chunk(100, b"small", 0).unwrap();
        // Invalid due to repeated index.
        buffer.add_chunk(100, b"small", 0).unwrap_err();

        let mut buffer = empty_buffer();
        buffer
            .add_chunk(10, b"Much bigger than ten characters", 0)
            .unwrap_err();

        let mut buffer = empty_buffer();
        buffer.add_chunk(29, b"Medium sized", 1).unwrap();
        // Cannot be smaller: only the terminator can (must come at the end.)
        buffer.add_chunk(29, b"small", 0).unwrap_err();
        buffer.add_chunk(29, b"small", 2).unwrap();
    }
}
