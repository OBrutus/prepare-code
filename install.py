from sys import platform
import os
import sys

TARGET_INSTALL_DIR = "~/.prepare-code"
# TARGET_INSTALL_DIR = "/tmp/prepare-code"

expressInstall = False
ConfigFileName = "config.yaml"


def remove_dir_if_exists(path, prompt=None):
    if os.path.isdir(os.path.expanduser(path)):
        print(f"Already existing directory at {path}!")
        if expressInstall or not prompt:
            choice = 'y'
        else:
            choice = input(prompt).strip().lower() or 'y'
        if choice == 'y':
            os.system(f"rm -rf {os.path.expanduser(path)}")
            return True
        else:
            print("Exiting setup.")
            return False
    return True


def clone_and_build():
    if os.path.isdir(os.path.expanduser("/tmp/prepare-code")):
        os.system("rm -rf /tmp/prepare-code")

    tmp_cmd = "cd /tmp && git clone https://github.com/OBrutus/prepare-code.git"
    # tmp_cmd = "cd /tmp && cp -r /Users/obrutus/kode/prepare-code /tmp/"
    
    status_code = os.system(tmp_cmd)

    if status_code != 0:
        print("Error during cloning repository. "
              "Please check your internet connection and try again.")
        return False

    status_code = os.system("cd /tmp/prepare-code && chmod +x build.sh && ./build.sh")

    if status_code != 0:
        print("Error during building the project. "
              "Please ensure you have Go installed and try again.")
        return False
    return True


def add_to_path():
    shell_config = os.path.expanduser("~/.bashrc")
    if platform == "darwin":
        shell_config = os.path.expanduser("~/.zshrc")
    with open(shell_config, "a") as f:
        f.write('\nexport PATH=' + TARGET_INSTALL_DIR + '":$PATH"\n')
    print(f"Added to PATH in {shell_config}. Please restart your terminal "
          "or run 'source {shell_config}' to apply changes.")


def setup_nix():
    if not remove_dir_if_exists(TARGET_INSTALL_DIR, "Do you want to remove it "
                                "and continue? (Y/n)"):
        return None

    if not clone_and_build():
        return None

    os.system("mkdir -p "+TARGET_INSTALL_DIR)
    os.system("mv /tmp/prepare-code/* " + TARGET_INSTALL_DIR)
    print("Setup completed to clone.")
    if expressInstall:
        choice = 'y'
    else:
        choice_prompt = "Do you want to add to PATH? (Y/n)"
        choice = input(choice_prompt).strip().lower() or 'y'
    if choice == 'y':
        add_to_path()
    return os.path.expanduser(TARGET_INSTALL_DIR)


def setup():
    if platform == "linux" or platform == "linux2":
        print("Setting up based upon linux platform")
        return setup_nix()
    elif platform == "darwin":
        # OS X
        print("Setting up based upon macOS platform")
        return setup_nix()
    elif platform == "win32":
        # Windows...
        print("Setting up based upon Windows platform")
        print("[-] Sorry! Currently not supported.")
        return None


def create_config_file(install_dir: str):
    content = "install_dir: " + install_dir + "\n"
    content += "bypass_prompt: " + "false" + "\n"
    content += "preferred_editor: nvim \n"
    content += "open_in_editor: false \n"

    file = open(os.path.join(install_dir, ConfigFileName), 'w')
    file.write(content)


def main():
    install_dir = setup()
    if install_dir is None:
        return

    create_config_file(install_dir)


if __name__ == "__main__":
    expressInstall = True if sys.argv[0] == '' else False

    print("This is install script")
    main()

