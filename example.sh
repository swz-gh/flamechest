export FLAMECHEST_AUTH_TOKEN="MY_AUTH_TOKEN"
export FLAMECHEST_ENDPOINT="https://example.com/file/upload"
flameshot gui -r | flamechest | xclip -selection c