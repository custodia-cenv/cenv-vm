import os
import subprocess
import platform

def build_go_program():
    # Paths
    src_dir = "./src/cmd"
    build_dir = "build"
    
    # Ensure the build directory exists
    if not os.path.exists(build_dir):
        os.makedirs(build_dir)
    
    # Determine architecture and OS
    arch = platform.machine()
    system_os = platform.system().lower()  # Use lowercase for consistency
    
    # Adjust Go architecture/OS naming convention if needed
    if arch == "x86_64":
        arch = "amd64"
    elif arch in ["armv7l", "armv6l"]:
        arch = "arm"
    elif arch == "aarch64":
        arch = "arm64"
    
    # Target binary name
    binary_name = f"cusxvm_{system_os}_{arch}"
    output_binary = os.path.join(build_dir, binary_name)
    
    # Go build command
    go_build_cmd = [
        "go", "build",
        "-o", output_binary,
        src_dir
    ]
    
    try:
        # Run the build command
        print(f"Building Go program from {src_dir} into {output_binary}...")
        subprocess.run(go_build_cmd, check=True)
        print(f"Build completed successfully! Output: {output_binary}")
    except subprocess.CalledProcessError as e:
        print("Build failed:", e)
    except FileNotFoundError:
        print("Go is not installed or not in PATH. Please check your installation.")

if __name__ == "__main__":
    build_go_program()
