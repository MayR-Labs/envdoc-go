## EnvDoc

### Overview



### Commands

- `envdoc create-example [file] [output defaults to .env.example]` : Generates an example file based on the environment variable keys found in the specified file. The values in the example file are set to empty strings. if output file already exists, the user is prompted to confirm overwriting it.

- `envdoc create-schema [file] [output defaults to .env.schema.json]` : Generates a JSON schema file based on the environment variable keys found in the specified file. The schema defines each key as a string type. if output file already exists, the user is prompted to confirm overwriting it.

- `envdoc arrange [file]` : Arrange and group environment variables keys in the specified file. Grouping means similar prefixes will be clustered together.

- `envdoc audit [file]` : Generates an extensive markdown report of missing environment keys and duplicated keys in the specified file. After the report is generated, the user is prompted to select: a. Show on CLI b. Copy report content c. Save to file (if user selects this he is prompted next for a file path, defaults to envdoc-audit-{timestamp}.md)

- `envdoc compare [file1] [file2] [fileN]` : Generates an extensive markdown report of keys that are missing across multiple specified files. After the report is generated, the user is prompted to select: a. Show on CLI b. Copy report content c. Save to file (if user selects this he is prompted next for a file path, defaults to envdoc-compare-{timestamp}.md)

- `envdoc sync [file1] [file2] [fileN]` : Synchronizes environment variable keys across multiple specified files. Missing keys in each file are added with empty values. Before making any changes, the user is prompted to confirm the synchronization action.

- `envdoc base64 [encode|decode] [file]` : Encodes or decodes the specified file using base64 encoding. The user is prompted to enter a preferred output file path (defaults to {originalfilename}.b64 for encoding and {originalfilename}.decoded for decoding). The processed file is then saved to the output file.

- `envdoc hash [file]` : Generates a SHA256 hash of the specified file's contents and displays it on the CLI. Also prompts the user to copy the hash to clipboard.

- `envdoc encrypt [file]` : The user is prompted to enter a password and a preferred output file path (defaults to {originalfilename}.encrypted). The specified file is then encrypted and saved to the output file.

- `envdoc decrypt [file]` : The user is prompted to enter the password used during encryption and a preferred output file path (defaults to the file name without the .encrypted extension). The specified encrypted file is then decrypted and saved to the output file.

- `envdoc to [json|yaml] [file]` : Converts the specified file to the desired format (JSON, YAML, or .env). The user is prompted to enter a preferred output file path (defaults to {originalfilename}.{extension}). The converted file is then saved to the output file.

- `envdoc from [json or yaml file]` : Converts the specified JSON or YAML file to a .env format. The user is prompted to enter a preferred output file path (defaults to {originalfilename}.env). The converted file is then saved to the output file.

- `envdoc validate [file] [schema-file]` : Validates the specified file against the provided JSON schema file. A report is generated detailing any discrepancies found during validation. After the report is generated, the user is prompted to select: a. Show on CLI b. Copy report content c. Save to file (if user selects this he is prompted next for a file path, defaults to envdoc-validate-{timestamp}.md)

- `envdoc doctor` : This audits and compares every .env file (.env, .env.*) except the ones with .encrypted extension in the current working directory. A comprehensive markdown report is generated detailing missing and duplicated keys across these files. After the report is generated, the user is prompted to select: a. Show on CLI b. Copy report content c. Save to file (if user selects this he is prompted next for a file path, defaults to envdoc-doctor-{timestamp}.md)

- `envdoc engineer` : This commands sync and arranges every .env file (.env, .env.*) except the ones with .encrypted extension in the current working directory. Before making any changes, the user is prompted to confirm the synchronization and arrangement action.

- `envdoc help` : Displays help information about the envdoc tool and its commands.

- `envdoc version` : Displays the current version of the envdoc tool.

- `envdoc documentation` : Opens the documentation for envdoc in the default web browser.

- `envdoc license` : Displays the license information for the envdoc tool.

- `envdoc changelog` : Displays the changelog for the envdoc tool.

- `envdoc authors` : Displays the authors of the envdoc tool "Built with ❤️ by MayR Labs \n <https://github.com/MayR-Labs>".

### Notes

- Use github.com/AlecAivazis/survey/v2 for interactive prompts
- Confirming action uses a pin code mechanism for safety. A random 6 digit pin is generated and shown to the user, who must then input the same pin to confirm the action.
- Encryption and decryption use aes-256-cbc, pbkdf2
- The tool should be very interactive that is when parameters are missing it prompts the user for input
- Extensive markdown reports should have a table of contents and be well formatted

