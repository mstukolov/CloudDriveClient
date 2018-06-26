# CloudDriveClient

Utility for working with cloud storage systems Dropbox, Google Drive, Yandex Drive. Works under Linux operating system. 

GET Utility
git clone https://github.com/mstukolov/CloudDriveClient.git

Build and Run Utility:
1. go build cli-tools.go
2. cli-tools -service=cloud provider(Example, dropbox, google, yandex)
  
DROPBOX API:
1. get file shared link:
https://www.dropbox.com/s/fileID/FileName?dl=0
1. cli-tools -service=dropbox
2. in terminal window will show next message: This Dropbox download utility
3. Enter dropbox fileID: <fileID>
4. Enter file name: <file name>(or any other...)
