# clean hosts file 
function Clean-File {
    param (
        [string]$filePath,
        [string[]]$extraStrings = @()  # 默认值为空数组
    )

    # Check if 'sed' is installed
    # scoop install sed
    if (-not (Get-Command sed -ErrorAction SilentlyContinue)) {
        Write-Host "'sed' command is not available. Please install 'sed' to use this function." -ForegroundColor Red
        return
    }

    Write-Host "Cleaning up $filePath..."

    # Perform the default cleanup actions
    sed -i 's/0\.0\.0\.0//g' $filePath
    sed -i '/#/d' $filePath
    sed -i '/^$/d' $filePath
    sed -i 's/ //g' $filePath

    # Add extra strings to the beginning of each line
    foreach ($str in $extraStrings) {
        Write-Host "Adding extra string '$str' to the beginning of each line in $filePath"
        sed -i "1i $str" $filePath
    }

    Write-Host "File cleanup completed!" -ForegroundColor Green
}


