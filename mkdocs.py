#!/usr/bin/env python3

import subprocess
import sys
from pathlib import Path
import shutil

# Key: Path to the export.py script
# Value: Destination Path where documentation should be exported
core_scripts = {
    "./teachyourselfcs/programming/Structure and Interpretation of Computer Programs/export.py": Path(
        "./docs/teachyourselfcs/programming/Structure and Interpretation of Computer Programs/"
    ),
}

def wipe_destination(destination_path: Path):
    """
    Deletes the destination directory if it exists to ensure a clean slate.
    """
    if destination_path.exists():
        print(f"Wiping existing directory: {destination_path}")
        shutil.rmtree(destination_path)
    else:
        print(f"Destination directory does not exist. No need to wipe: {destination_path}")
    
    # Recreate the empty destination directory
    destination_path.mkdir(parents=True, exist_ok=True)
    print(f"Created fresh directory: {destination_path}")

def execute_export_script(export_script: Path, destination_path: Path):
    """
    Executes the export.py script with the destination path as an argument.
    """
    if not export_script.is_file():
        print(f"Error: Export script not found at {export_script}", file=sys.stderr)
        sys.exit(1)
    
    print(f"\nExecuting export script: {export_script}")
    print(f"Destination path: {destination_path}")
    
    try:
        subprocess.run([sys.executable, str(export_script), str(destination_path)], check=True)
        print(f"Successfully executed {export_script}")
    except subprocess.CalledProcessError as e:
        print(f"Error executing {export_script}: {e}", file=sys.stderr)
        sys.exit(1)

def main():
    """
    Main function to wipe target directories and execute export scripts.
    """
    print("Starting export process...")

    for export_script_str, destination_path in core_scripts.items():
        export_script = Path(export_script_str).resolve()
        destination_path = destination_path.resolve()
        
        print(f"\nProcessing export for: {export_script.name}")
        print(f"Target Destination: {destination_path}")
        
        # Step 1: Wipe the destination directory
        wipe_destination(destination_path)
        
        # Step 2: Execute the export script
        execute_export_script(export_script, destination_path)
    
    print("\nAll export processes completed successfully.")

if __name__ == '__main__':
    main()

