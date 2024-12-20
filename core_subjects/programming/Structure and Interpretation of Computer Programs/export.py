#!/usr/bin/env python3

import argparse
from pathlib import Path
import shutil
import subprocess
import sys


def execute_make_docs(script_dir):
    """
    Executes the make-docs.py script located in the script_dir.
    """
    make_docs_script = script_dir / "mk-docs.py"
    if not make_docs_script.exists():
        print(f"Error: {make_docs_script} does not exist.", file=sys.stderr)
        sys.exit(1)

    print(f"Executing {make_docs_script}...")
    try:
        subprocess.run([sys.executable, str(make_docs_script)], check=True)
        print("make-docs.py executed successfully.")
    except subprocess.CalledProcessError as e:
        print(f"Error executing make-docs.py: {e}", file=sys.stderr)
        sys.exit(1)


def copy_markdown_files(src_dir, dst_dir):
    """
    Copies all .md files from src_dir to dst_dir.
    """
    if not src_dir.exists():
        print(
            f"Warning: Source directory {src_dir} does not exist. Skipping.",
            file=sys.stderr,
        )
        return

    dst_dir.mkdir(parents=True, exist_ok=True)
    md_files = list(src_dir.glob("*.md"))

    if not md_files:
        print(f"No .md files found in {src_dir}.")
        return

    for md_file in md_files:
        shutil.copy(md_file, dst_dir)
        print(f"Copied {md_file} to {dst_dir}")


def copy_readme(src_dir, dst_dir):
    """
    Copies README.md from src_dir to dst_dir as index.md.
    """
    readme_src = src_dir / "README.md"
    index_dst = dst_dir / "index.md"

    if not readme_src.exists():
        print(f"Warning: {readme_src} does not exist. Skipping.", file=sys.stderr)
        return

    shutil.copy(readme_src, index_dst)
    print(f"Copied {readme_src} to {index_dst}")


def main():
    # Set up argument parser
    parser = argparse.ArgumentParser(description="Create a documentation structure.")
    parser.add_argument(
        "path", type=Path, help="Destination Path where the structure will be created."
    )
    args = parser.parse_args()
    destination_path = args.path.resolve()

    print(f"Destination Path: {destination_path}")

    # Get the directory where this script is located
    script_dir = Path(__file__).parent.resolve()

    # Step 1: Execute make-docs.py
    execute_make_docs(script_dir)

    # Step 2: Copy all .md files from ./exercises to p/Exercises
    exercises_src = script_dir / "exercises"
    exercises_dst = destination_path / "Exercises"
    copy_markdown_files(exercises_src, exercises_dst)

    # Step 3: Copy all .md files from ./notes to p/notes
    notes_src = script_dir / "notes"
    notes_dst = destination_path / "notes"
    copy_markdown_files(notes_src, notes_dst)

    # Step 4: Copy README.md to p/index.md
    copy_readme(script_dir, destination_path)

    print("Documentation structure created successfully.")


if __name__ == "__main__":
    main()
