# Run build scripts using the following:
Enter the .csproj file, add the following code:
```
<Target Name="" BeforeTargets="Build">
  <Exec Command="powershell" Arguments="-File $(ProjectDir)myscript.ps1"/>
</Target>
```
This XML will run the Powershell command in the project directory.

Check the following: https://learn.microsoft.com/en-us/aspnet/web-forms/overview/deployment/advanced-enterprise-web-deployment/running-windows-powershell-scripts-from-msbuild-project-files
