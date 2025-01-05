set VTAG=v1.4.6

git add -A
git commit -m "%VTAG% commit"
git push

git tag -a "%VTAG%" -m "version tag"
git push origin "%VTAG%"
