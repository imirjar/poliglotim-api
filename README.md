# Docker Image Publishing

This project automatically builds and publishes Docker images to GitHub Container Registry (GHCR) when changes are pushed to the main branch.

## 📦 Published Images

The workflow publishes two tags for each image:
- `ghcr.io/owner/image-name:latest` (latest stable version)
- `ghcr.io/owner/image-name:commit-sha` (specific commit version)

## 🔧 Usage

### Pull the Docker Image

```bash
docker pull ghcr.io/your-username/your-image-name:latest
```

### Run the Container

```bash
docker run -d ghcr.io/your-username/your-image-name:latest
```

## 📋 Requirements

- GitHub repository with enabled GitHub Actions
- Proper workflow permissions (read and write)
- Dockerfile in the repository root

## 🚀 Automated Workflow

The build process is automatically triggered on:
- ✅ Pushes to the `main` branch
- ✅ Pull requests to the `main` branch  
- ✅ Manual triggers via GitHub Actions UI

## 🔒 Authentication

When pulling images from private repositories, you need to authenticate:

```bash
echo ${{ secrets.GITHUB_TOKEN }} | docker login ghcr.io -u YOUR_USERNAME --password-stdin
```

## 📁 Image Location

Published images are available in:
- **GitHub Packages**: `https://github.com/YOUR_USERNAME?tab=packages`
- **Container Registry**: `ghcr.io/YOUR_USERNAME/IMAGE_NAME`

## ⚙️ Configuration

The following environment variables are used in the workflow:

| Variable | Description |
|----------|-------------|
| `DOCKER_IMAGE_NAME` | Name of the Docker image |
| `DOCKER_CONTAINER_NAME` | Container name for running instances |
| `APP_FILE` | Main application file |
| `BIN_NAME` | Binary output name |

## 🔄 Manual Trigger

You can manually trigger the build process:
1. Go to **Actions** tab in your repository
2. Select **"Build image after tests"** workflow
3. Click **"Run workflow"**

## 📝 Notes

- Images are built using Docker Buildx for better performance
- All images are automatically scanned for vulnerabilities by GitHub
- The `latest` tag always points to the most recent successful build
- Commit-specific tags provide immutable references for deployment