import os
import re
from pathlib import Path

# Define the root directory (where this script is located)
ROOT_DIR = Path(__file__).parent.resolve()
EXERCISES_DIR = ROOT_DIR / "exercises"
OUTPUT_DIR = ROOT_DIR / "exercises"

# SICP section to file number mapping
SECTION_FILE_MAP = {
    "1.1": "10",
    "1.2": "11",
    "1.3": "12",
    "2.1": "14",
    "2.2": "15",
    "2.3": "16",
    "2.4": "17",
    "2.5": "18",
    "3.1": "20",
    "3.2": "21",
    "3.3": "22",
    "3.4": "23",
    "3.5": "24",
    "4.1": "26",
    "4.2": "27",
    "4.3": "28",
    "4.4": "29",
    "5.1": "31",
    "5.2": "32",
    "5.3": "33",
    "5.4": "34",
    "5.5": "35"
}

def get_file_number(section):
    """Get the correct file number for a section"""
    # Extract main section number (e.g., "1.1" from "1.1.6")
    main_section = '.'.join(section.split('.')[:2])
    return SECTION_FILE_MAP.get(main_section, "4")  # Default to 4 if not found

def generate_section_link(section):
    """Generate a link to the section in the SICP online book"""
    section_match = re.match(r"(\d+\.\d+\.?\d*)", section)
    if section_match:
        section_num = section_match.group(1)
        file_num = get_file_number(section_num)
        base_url = f"https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-{file_num}.html"
        return f"{base_url}#%_sec_{section_num}"
    else:
        raise ValueError(f"Invalid section title: {section}")

def generate_exercise_link(exercise_num, section_title):
    """Generate a link to the exercise in the SICP online book"""
    file_num = get_file_number(section_title)
    base_url = f"https://mitp-content-server.mit.edu/books/content/sectbyfn/books_pres_0/6515/sicp.zip/full-text/book/book-Z-H-{file_num}.html"
    return f"{base_url}#%_thm_{exercise_num}"

def handle_katex_in_comment(line):
    """Preserve KaTeX expressions in comments by ensuring proper spacing"""
    line = re.sub(r'(\$\$?)([^\$]+)(\$\$?)', lambda m: 
                  ('\n\n' + m.group(0) + '\n\n') if m.group(1) == '$$' else m.group(0), 
                  line)
    return line

def parse_scm_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        lines = file.readlines()

    section_title = ""
    exercise_number = ""
    content_segments = []
    current_block = []
    metadata_lines = []
    in_metadata = True

    for line in lines:
        stripped_line = line.strip()
        
        # Handle metadata section
        if in_metadata:
            if stripped_line.startswith(";; Section"):
                section_title = stripped_line.replace(";; Section", "").strip()
                metadata_lines.append(line)
            elif stripped_line.startswith(";; Exercise"):
                exercise_match = re.search(r";;\s*Exercise\s+([\d\.]+)", stripped_line)
                if exercise_match:
                    exercise_number = exercise_match.group(1)
                metadata_lines.append(line)
            elif stripped_line.startswith(";;"):
                metadata_lines.append(line)
            elif stripped_line == "":
                metadata_lines.append(line)
            else:
                in_metadata = False
                current_block.append(line)
        else:
            if stripped_line.startswith(";;md"):
                if current_block:
                    content_segments.append(("code", "".join(current_block)))
                    current_block = []
                
                md_content = stripped_line.replace(";;md", "").strip()
                md_content = handle_katex_in_comment(md_content)
                content_segments.append(("markdown", md_content))
            else:
                current_block.append(line)

    # Add any remaining code block
    if current_block:
        content_segments.append(("code", "".join(current_block)))

    return {
        "section_title": section_title,
        "section_link": generate_section_link(section_title),
        "exercise_number": exercise_number,
        "exercise_link": generate_exercise_link(exercise_number, section_title),
        "metadata": "".join(metadata_lines),
        "content_segments": content_segments
    }

def process_chapter(chapter_path):
    chapter_name = chapter_path.name
    chapter_number_match = re.match(r"chapter-(\d+)", chapter_name, re.IGNORECASE)
    if not chapter_number_match:
        print(f"Skipping unrecognized directory: {chapter_name}")
        return None

    chapter_number = chapter_number_match.group(1)
    markdown_filename = OUTPUT_DIR / f"chapter-{chapter_number}.md"
    print(f"Processing {chapter_name} -> {markdown_filename.name}")

    scm_files = sorted(
        chapter_path.glob("*.scm"),
        key=lambda x: float(re.search(r"(\d+\.\d+)", x.stem).group(1)) if re.search(r"(\d+\.\d+)", x.stem) else 0
    )

    if not scm_files:
        print(f"No .scm files found in {chapter_name}. Skipping.")
        return None

    markdown_content = f"# Chapter {chapter_number}\n\n"
    current_section = None

    for scm_file in scm_files:
        data = parse_scm_file(scm_file)

        if data["section_title"] and data["section_title"] != current_section:
            markdown_content += f"## [Section {data['section_title']}]({data['section_link']})\n\n"
            current_section = data["section_title"]

        markdown_content += f"### [Exercise {data['exercise_number']}]({data['exercise_link']})\n\n"

        for segment_type, content in data["content_segments"]:
            if segment_type == "code":
                markdown_content += "```scheme\n"
                markdown_content += content.rstrip() + "\n"
                markdown_content += "```\n\n"
            else:
                markdown_content += content + "\n\n"

    with open(markdown_filename, 'w', encoding='utf-8') as md_file:
        md_file.write(markdown_content)

    print(f"Generated {markdown_filename.name}")
    return markdown_filename.name

def make_doc():
    if not EXERCISES_DIR.exists() or not EXERCISES_DIR.is_dir():
        print(f"Exercises directory not found at {EXERCISES_DIR}")
        return

    chapter_dirs = sorted(
        [d for d in EXERCISES_DIR.iterdir() if d.is_dir() and re.match(r"chapter-\d+", d.name, re.IGNORECASE)],
        key=lambda x: int(re.match(r"chapter-(\d+)", x.name, re.IGNORECASE).group(1))
    )

    if not chapter_dirs:
        print("No chapter directories found.")
        return

    for chapter_dir in chapter_dirs:
        process_chapter(chapter_dir)

if __name__ == "__main__":
    make_doc()
