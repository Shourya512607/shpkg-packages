# shpkg.py
import argparse, os, subprocess

REPO = "./repo"

def install(pkg):
    path = os.path.join(REPO, pkg, "install.sh")
    if os.path.exists(path):
        subprocess.run(["sh", path])
        print(f"✅ {pkg} installed.")
    else:
        print(f"❌ Package {pkg} not found.")

def list_packages():
    print("📦 Available packages:")
    for p in os.listdir(REPO):
        print(" -", p)

parser = argparse.ArgumentParser(description="SHPKG Package Manager")
parser.add_argument("command", help="install, list")
parser.add_argument("package", nargs='?', help="package name")
args = parser.parse_args()

if args.command == "install" and args.package:
    install(args.package)
elif args.command == "list":
    list_packages()
else:
    print("❓ Usage: python shpkg.py [install <pkg>] | [list]")
