# github-actions-cron-readme-weather-api

# 🚀 Update your README with the current weather forcast using Go Templates, Docker, GitHub Actions 🚀

https://github.com/coding-to-music/github-actions-cron-readme-weather-api

From / By https://github.com/huantt/weather-forecast

https://dev.to/jacktt/creating-dynamic-readmemd-file-388o

https://github.com/FahimFBA/test

<!-- <div style="text-align:center;">
  <img src="/images/chakra.jpg" alt="Image" />
  <p><em>Chakra Component Library with Next.js</em></p>
</div> -->

## Go Environment:

```java
Go Version
```

Output

```java
go version go1.17.6 linux/amd64
```

## Environment variables:

```java
none
```

## GitHub

```java
git init
git add .
git remote remove origin
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:coding-to-music/github-actions-cron-readme-weather-api.git
git push -u origin main
```

## signup for a free Weatherapi.com account

https://www.weatherapi.com/signup.aspx

Put Repo secrets WEATHER_API_KEY={weatherapi.com API secret goes here}

## Install latest Go version

Latest downloads are available here:

https://go.dev/dl/

```java
wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz
```

### Extract the Archive:

Use the tar command to extract the contents of the downloaded archive. Replace "your-username" with your actual username:

```java
sudo tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz
```

This command will extract the Go files into the /usr/local directory.

### Set Go Environment Variables:

You need to set Go's environment variables so that your system knows where to find the Go binaries. Add the following lines to your shell profile file (e.g., ~/.bashrc, ~/.zshrc, or ~/.profile), replacing "your-username" with your actual username if necessary:

```java
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

### After adding these lines, run the following command to apply the changes immediately to your current session:

```java
source ~/.bashrc  # or ~/.zshrc, or the appropriate file for your shell
```

### Verify the Installation:

After setting the environment variables, you can verify the installation by running:

```java
go version
```

It should display the Go version you downloaded and installed.

That's it! You've successfully installed Go from the tar.gz archive on your Linux system. You can now start using Go for your development projects.

## To create a release for your GitHub repository and assign a version number to it, follow these steps:

Navigate to Your Repository: Go to the GitHub repository where you want to create a release.

Click on the "Releases" Tab: On the repository's main page, click on the "Releases" tab, usually located between "Code" and "Packages."

Create a New Release:

a. Click on the "Draft a new release" button or "Create a new release" if you don't see the "Draft a new release" button.

b. Fill in the release information:

Tag version: This is where you specify the version number. It's common to use semantic versioning (e.g., "v1.0.0"). Make sure it matches the version your actions are looking for.

Release title: You can give your release a descriptive title.

Description: Provide release notes or any relevant information about the changes in this release.

Attach binaries (optional): If you have any files, such as compiled binaries or artifacts, that you want to associate with this release, you can attach them here.

c. When you're ready, click the "Publish release" button.

Publish the Release: Once you publish the release, it will be available for others to see in the "Releases" section of your repository.

Update Your Workflow: If your GitHub Actions workflow is looking for a specific version, you will need to update the workflow configuration (usually in a YAML file) to reference the new release version. Make sure the workflow uses the correct tag/version number you just created.

Commit and Push Changes: After updating your workflow file, commit the changes to your repository and push them to GitHub. This will trigger your workflow with the new release version.

Now, your GitHub Actions workflow should be able to use the specific release version you've created. Remember to update the version reference in your workflow whenever you create a new release with a new version number.

## Modify two files to utilize the specific package number each time a new release is created

```java
.github/workflows/run.yml

README.md.template
```

## Changes to the template for F temp

Removed C temp and replaced with F temp

File `handler/collector/template/daily-forcast.md.template`

```java

# Temperature
            <td>{{ $weather.MinTempC }} -  {{ $weather.MaxTempC}} °C</td>

# Wind
            <td>{{ $weather.AvgWindKph}} kph</td>
```

File `handler/collector/template/hourly-forcast.md.template`

```java

# Temperature
            <td>{{ $weather.AvgTempC}} °C</td>

# Wind
            <td>{{ $weather.AvgWindKph}} kph</td>
```
