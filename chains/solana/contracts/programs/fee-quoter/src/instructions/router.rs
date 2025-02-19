use crate::state::CodeVersion;

use super::interfaces::*;
use super::v1;

/**
 * This file routes traffic between multiple versions of our business logic, which can be upgraded in a
 * backwards-compatible way and so we can gradually shift traffic between versions (and rollback if there are issues).
 * This also supports flexible criteria on how to shift traffic between versions (e.g. per-lane, all at once, etc)
 * for each module (e.g. the criteria for admin actions may differ the criteria for public-facing interfaces).
 *
 * On any code changes to the business logic, a new version of the module should be created and leave the previous one
 * untouched. The new version should be added to the match statement below so that traffic can be progressively shifted,
 * and is possible to rollback easily via config changes without having to redeploy.
 *
 * As we currently have a single version, all branches lead to the same outcome, but the code is structured in a way
 * that is easy to extend to multiple versions.
 */

pub fn public(
    lane_code_version: CodeVersion,
    default_code_version: CodeVersion,
) -> &'static dyn Public {
    // The lane-specific code version takes precedence over the default code version.
    // If the lane just specifies using the default, then we use that one.
    match lane_code_version {
        CodeVersion::V1 => &v1::public::Impl,
        CodeVersion::Default => match default_code_version {
            CodeVersion::Default => &v1::public::Impl, // can't happen, but default to v1 so the `match` is exhaustive
            CodeVersion::V1 => &v1::public::Impl,
        },
    }
}

pub fn prices(code_version: CodeVersion) -> &'static dyn Prices {
    match code_version {
        CodeVersion::Default => &v1::prices::Impl,
        CodeVersion::V1 => &v1::prices::Impl,
    }
}

pub fn admin(code_version: CodeVersion) -> &'static dyn Admin {
    match code_version {
        CodeVersion::Default => &v1::admin::Impl,
        CodeVersion::V1 => &v1::admin::Impl,
    }
}
