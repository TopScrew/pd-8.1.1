start=$(date +%s)
/Users/screw/Desktop/snyk-macos-arm64 test --package-manager=gomodules --severity-threshold=low
end=$(date +%s)
echo "Execution time: $((end - start)) seconds"
