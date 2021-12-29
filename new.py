import sys

def main():
    title = sys.argv[1]
    create_directory_with_given_name(title)


def create_directory_with_given_name(name: str):
    import os

    os.mkdir(name)
    # create a markdown file with the same name inside the directory
    with open(name + "/" + name + ".md", "w") as f:
        f.write(f"# [{name}]()\n")
    # create a go file with the same name inside the directory
    with open(name + "/" + name + ".go", "w") as f:
        f.write(f"package main\n\n")


if __name__ == "__main__":
    main()
