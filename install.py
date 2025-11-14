from sys import platform
import os
import sys


expressInstall = False

def setupNix():
    if os.path.isdir(os.path.expanduser("~/.prepare-code")):
        print("Already existing directory!")
        choice = 'y' if expressInstall else (input("Do you want to remove it and continue? (Y/n)").strip().lower() or 'y')
        if choice == 'y':
            os.system("rm -rf ~/.prepare-code")
        else:
            print("Exiting setup.")
            return
        setupNix()
        return

    # now we have new canvas
    if os.path.isdir(os.path.expanduser("/tmp/prepare-code")):
        os.system("rm -rf /tmp/prepare-code")

    statusCode = os.system("cd /tmp && "
                           # "git clone https://github.com/OBrutus/prepare-code.git")
                           "cp -r /Users/obrutus/kode/compCode/prepare-code /tmp/prepare-code")
    if statusCode != 0:
        print("Error during cloning repository. "
              "Please check your internet connection and try again.")
        return

    # post clone install the repo
    statusCode = os.system(
        "cd /tmp/prepare-code && chmod +x build.sh && ./build.sh"
    )
    if statusCode != 0:
        print("Error during building the project. "
              "Please ensure you have Go installed and try again.")
        return

    # now cloned repo is in /tmp/prepare-code
    os.system("mkdir -p ~/.prepare-code")
    os.system("mv /tmp/prepare-code/* ~/.prepare-code")

    print("Setup completed to clone.")
    choice = 'y' if expressInstall else (input("Do you want to add to PATH? (Y/n)").strip().lower() or 'y')
    if choice == 'y':
        shellConfig = os.path.expanduser("~/.bashrc")
        if platform == "darwin":
            shellConfig = os.path.expanduser("~/.zshrc")

        with open(shellConfig, "a") as f:
            f.write('\nexport PATH="$HOME/.prepare-code:$PATH"\n')

        print(f"Added to PATH in {shellConfig}. Please restart your terminal or run 'source {shellConfig}' to apply changes.")


def main():
    if platform == "linux" or platform == "linux2":
        print("Setting up based upon linux platform")
        setupNix()
    elif platform == "darwin":
        # OS X
        print("Setting up based upon macOS platform")
        setupNix()
    elif platform == "win32":
        # Windows...
        print("Setting up based upon Windows platform")
        print("[-] Sorry! Currently not supported.")


if __name__ == "__main__":
    expressInstall = True if sys.argv[0] == '' else False

    print("This is install script")
    main()

