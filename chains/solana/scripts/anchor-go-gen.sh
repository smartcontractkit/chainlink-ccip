#!/usr/bin/env bash

set -e

# Delete all .go files (anchor-go v1.0.0 generates all needed types directly)
find ./gobindings/latest -name "*.go" -type f -delete

# anchor-go internally runs "go fmt" and "go build" on generated packages and panics
# on failure. This causes issues when generated code has name collisions (e.g. fee_quoter)
# that we fix in post-processing. Create a wrapper that no-ops these commands so anchor-go
# completes, then we fix the code and run go fmt ourselves at the end.
REAL_GO=$(which go)
GO_WRAPPER_DIR=$(mktemp -d)
trap 'rm -rf "$GO_WRAPPER_DIR"' EXIT
cat > "$GO_WRAPPER_DIR/go" << GOWRAPPER
#!/bin/bash
if [ "\$1" = "fmt" ] || [ "\$1" = "build" ]; then
  exit 0
fi
exec "$REAL_GO" "\$@"
GOWRAPPER
chmod +x "$GO_WRAPPER_DIR/go"

function generate_bindings() {
  local idl_path_str="$1"
  IFS='/' read -r -a idl_path <<< "${idl_path_str}"
  IFS='.' read -r -a idl_name <<< "${idl_path[3]}"
  PATH="$GO_WRAPPER_DIR:$PATH" anchor-go -idl "${idl_path_str}" -output ./gobindings/latest/"${idl_name}" -no-go-mod
}

# Generate bindings for all IDLs (including vendor)
for idl_path_str in "contracts/target/idl"/*
do
  generate_bindings "${idl_path_str}"
done

# Fix vendor IDL instruction discriminators: The vendor .so files were compiled with
# old Anchor (snake_case discriminators), but the vendor IDL JSON files use the new
# Anchor format (camelCase discriminators). Recompute instruction discriminators using
# snake_case naming to match the deployed .so binaries.
python3 - << 'PYEOF'
import hashlib, json, re, glob

def camel_to_snake(name):
    s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', name)
    return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()

for path in glob.glob('contracts/target/vendor/*.json'):
    with open(path) as f:
        idl = json.load(f)
    modified = False
    for ix in idl.get('instructions', []):
        snake = camel_to_snake(ix['name'])
        correct = list(hashlib.sha256(f'global:{snake}'.encode()).digest()[:8])
        if ix['discriminator'] != correct:
            ix['discriminator'] = correct
            modified = True
    if modified:
        with open(path, 'w') as f:
            json.dump(idl, f, indent=2)
            f.write('\n')
PYEOF

for idl_path_str in "contracts/target/vendor"/*.json
do
  generate_bindings "${idl_path_str}"
done

# Restore original vendor IDLs (we only modified them temporarily for generation)
git checkout -- contracts/target/vendor/*.json

# Fix anchor-go name collision in fee_quoter: "config" is both an instruction arg
# (BillingTokenConfig) and an account (PublicKey), both PascalCased to "Config".
# Rename the account field to "ConfigAccount" in the affected structs and methods.
FQ_INSTRUCTIONS="./gobindings/latest/fee_quoter/instructions.go"
if [ -f "$FQ_INSTRUCTIONS" ]; then
  python3 - "$FQ_INSTRUCTIONS" << 'PYEOF'
import sys

file_path = sys.argv[1]
with open(file_path, 'r') as f:
    content = f.read()

for struct in ['AddBillingTokenConfigInstruction', 'UpdateBillingTokenConfigInstruction']:
    # 1. Rename account field in struct declaration
    content = content.replace(
        'type ' + struct + ' struct {\n\tConfig BillingTokenConfig `json:"config"`\n\n\t// Accounts:\n\tConfig ',
        'type ' + struct + ' struct {\n\tConfig BillingTokenConfig `json:"config"`\n\n\t// Accounts:\n\tConfigAccount ',
    )

    # 2. Fix PopulateFromAccountIndices: obj.Config -> obj.ConfigAccount for the config account
    method_sig = 'func (obj *' + struct + ') PopulateFromAccountIndices'
    idx = content.find(method_sig)
    if idx >= 0:
        comment = '// Set config account from index'
        comment_idx = content.find(comment, idx)
        if comment_idx >= 0:
            old = 'obj.Config = accountKeys[indices[indexOffset]]'
            new = 'obj.ConfigAccount = accountKeys[indices[indexOffset]]'
            replace_idx = content.find(old, comment_idx)
            if replace_idx >= 0 and replace_idx < comment_idx + 300:
                content = content[:replace_idx] + new + content[replace_idx + len(old):]

    # 3. Fix GetAccountKeys: first append uses the config account field
    old_get = ('func (obj *' + struct + ') GetAccountKeys() []solanago.PublicKey {\n'
               '\tkeys := make([]solanago.PublicKey, 0)\n'
               '\tkeys = append(keys, obj.Config)')
    new_get = ('func (obj *' + struct + ') GetAccountKeys() []solanago.PublicKey {\n'
               '\tkeys := make([]solanago.PublicKey, 0)\n'
               '\tkeys = append(keys, obj.ConfigAccount)')
    content = content.replace(old_get, new_get)

with open(file_path, 'w') as f:
    f.write(content)
PYEOF
fi

# Fix anchor-go name collision in timelock: The IDL has instructions named "initialize"
# and "initializeInstruction" (and similarly "initializeBypasserInstruction"). The generator
# maps both to the same Go name (e.g. NewInitializeInstruction). Rename the *Instruction
# variants to append an extra "Instruction" suffix to disambiguate.
TL_INSTRUCTIONS="./gobindings/latest/timelock/instructions.go"
if [ -f "$TL_INSTRUCTIONS" ]; then
  python3 - "$TL_INSTRUCTIONS" << 'PYEOF'
import re, sys

file_path = sys.argv[1]
with open(file_path, 'r') as f:
    content = f.read()

# For each pair, rename the "*Instruction" variant (from "initializeInstruction" and
# "initializeBypasserInstruction" IDL names) by appending an extra "Instruction" suffix.
# The discriminator constants (Instruction_Xxx) must NOT be renamed.
renames = {
    'InitializeInstruction': '__DISC_PROTECT_INIT_IX__',
    'InitializeBypasserInstruction': '__DISC_PROTECT_INIT_BYP_IX__',
}
for name, placeholder in renames.items():
    # Protect discriminator constants from replacement (placeholder must NOT contain the target)
    content = content.replace('Instruction_' + name, placeholder)
    # Rename all remaining occurrences (struct, methods, constructors, string literals)
    # Use negative lookahead to avoid matching longer names (e.g. InitializeInstructionOperation)
    content = re.sub(re.escape(name) + r'(?!Operation)', name + 'Instruction', content)
    # Restore discriminator constants
    content = content.replace(placeholder, 'Instruction_' + name)

with open(file_path, 'w') as f:
    f.write(content)
PYEOF
fi

go fmt ./...
