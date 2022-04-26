Register-ArgumentCompleter -Native -CommandName goo -ScriptBlock {
    param($commandName, $wordToComplete, $cursorPosition)
    $other = "$wordToComplete --generate-bash-completion"
    Invoke-Expression $other | ForEach-Object {
        [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
    }
}
