import sys
import click
from typing import List

@click.command()
@click.option('--tags', '-t', help='题目标签，支持多个，使用英文逗号分隔')
@click.argument('title')
def main(title, tags):
    if tags:
        tags = tags.split(',')
    else:
        tags = []

    create_directory_with_given_name(title, tags)


def create_directory_with_given_name(name: str, tags: List[str]):
    import os

    os.mkdir(name)
    # create a markdown file with the same name inside the directory
    with open(name + "/" + name + ".md", "w") as f:
        content = f"# [{name}]()\n"
        if len(tags) > 0:
            content += "\n\n"
            content += ' '.join(f'`{tag}`' for tag in tags)
        f.write(content)

    # create a go file with the same name inside the directory
    with open(name + "/" + name + ".go", "w") as f:
        f.write(f"package main\n\n")


if __name__ == "__main__":
    main()
