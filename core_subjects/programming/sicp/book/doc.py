import os
import re
from pathlib import Path

# Define the root directory (where this script is located)
ROOT_DIR = Path(__file__).parent.resolve()

# Define paths
EXERCISES_DIR = ROOT_DIR / "exercises"
OUTPUT_DIR = ROOT_DIR  # Output Markdown files in the root directory
MASTER_README = ROOT_DIR / "README.md"

# Function to extract information from a .scm file
def parse_scm_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        lines = file.readlines()

    section_title = ""
    exercise_number = ""
    exercise_title = ""
    exercise_description = []
    content_segments = []  # List to hold tuples of (type, content)

    current_code = []
    in_metadata = True  # Flag to determine if we are still parsing metadata

    for line in lines:
        stripped_line = line.strip()

        if in_metadata:
            if stripped_line.startswith(";; Section"):
                section_title = stripped_line.replace(";; Section", "").strip()
                continue
            elif stripped_line.startswith(";; Exercise"):
                # Parse Exercise Number and Title
                exercise_num_match = re.match(r";; Exercise (\d+\.\d+):\s*(.*)", stripped_line)
                if exercise_num_match:
                    exercise_number = exercise_num_match.group(1)
                    exercise_title = exercise_num_match.group(2)
                continue
            elif stripped_line.startswith(";;"):
                # Assume description lines
                desc = stripped_line.replace(";;", "").strip()
                exercise_description.append(desc)
                continue
            elif stripped_line == "":
                # Blank line, possibly end of metadata
                continue
            else:
                # First code line detected
                in_metadata = False
                if not stripped_line.startswith(";;"):
                    current_code.append(line.rstrip())
                else:
                    # Line starts with `;;` but we're out of metadata; treat as code comment
                    current_code.append(line.rstrip())
        else:
            # In code mode
            if stripped_line.startswith(";;md"):
                # Split code block at `;;md` comment
                if current_code:
                    code = "\n".join(current_code)
                    content_segments.append(("code", code))
                    current_code = []
                # Add markdown comment
                md_comment = stripped_line.replace(";;md", "").strip()
                content_segments.append(("md", md_comment))
            else:
                # All other lines are part of code, including lines starting with `;;`
                current_code.append(line.rstrip())

    # Append any remaining code after processing all lines
    if current_code:
        code = "\n".join(current_code)
        content_segments.append(("code", code))

    return {
        "section_title": section_title,
        "exercise_number": exercise_number,
        "exercise_title": exercise_title,
        "exercise_description": " ".join(exercise_description),
        "content_segments": content_segments
    }

# Function to process each chapter
def process_chapter(chapter_path):
    chapter_name = chapter_path.name  # e.g., "chapter-1"
    chapter_number_match = re.match(r"chapter-(\d+)", chapter_name, re.IGNORECASE)
    if not chapter_number_match:
        print(f"Skipping unrecognized directory: {chapter_name}")
        return None

    chapter_number = chapter_number_match.group(1)
    markdown_filename = OUTPUT_DIR / f"chapter-{chapter_number}.md"
    print(f"Processing {chapter_name} -> {markdown_filename.name}")

    # Collect all .scm files and sort them based on their numeric prefix
    scm_files = sorted(
        chapter_path.glob("*.scm"),
        key=lambda x: int(x.stem.split('.')[1]) if len(x.stem.split('.')) > 1 and x.stem.split('.')[1].isdigit() else 0
    )

    if not scm_files:
        print(f"No .scm files found in {chapter_name}. Skipping.")
        return None

    # Initialize Markdown content
    markdown_content = f"# Chapter {chapter_number}\n\n"

    current_section = None  # To track the current section and avoid duplication

    for scm_file in scm_files:
        data = parse_scm_file(scm_file)

        # Add Section Title if it's new
        if data["section_title"]:
            if data["section_title"] != current_section:
                markdown_content += f"## Section {data['section_title']}\n\n"
                current_section = data["section_title"]

        # Add Exercise Header
        exercise_num_formatted = data["exercise_number"]  # Already in X.Y format
        exercise_title = data["exercise_title"]
        markdown_content += f"### Exercise {exercise_num_formatted}: {exercise_title}\n\n"

        # Add Exercise Description
        markdown_content += f"{data['exercise_description']}\n\n"

        # Add Content Segments
        for segment_type, content in data["content_segments"]:
            if segment_type == "code":
                markdown_content += "```scm\n"
                markdown_content += f"{content}\n"
                markdown_content += "```\n\n"
            elif segment_type == "md":
                markdown_content += f"{content}\n\n"

    # Write to the Markdown file
    with open(markdown_filename, 'w', encoding='utf-8') as md_file:
        md_file.write(markdown_content)

    print(f"Generated {markdown_filename.name}\n")
    return markdown_filename.name

# Main execution
def main():
    if not EXERCISES_DIR.exists() or not EXERCISES_DIR.is_dir():
        print(f"Exercises directory not found at {EXERCISES_DIR}")
        return

    # Find all chapter directories (e.g., chapter-1, chapter-2, ...)
    chapter_dirs = sorted(
        [d for d in EXERCISES_DIR.iterdir() if d.is_dir() and re.match(r"chapter-\d+", d.name, re.IGNORECASE)],
        key=lambda x: int(re.match(r"chapter-(\d+)", x.name, re.IGNORECASE).group(1))
    )

    if not chapter_dirs:
        print("No chapter directories found.")
        return

    generated_markdown_files = []

    for chapter_dir in chapter_dirs:
        markdown_file = process_chapter(chapter_dir)
        if markdown_file:
            generated_markdown_files.append(markdown_file)

    # Update the master README.md with links to all chapters
    if not generated_markdown_files:
        print("No Markdown files were generated.")

if __name__ == "__main__":
    main()

