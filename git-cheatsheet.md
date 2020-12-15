* Create new branch and check it out:

   ```
   git checkout -b my-new-branch
   
   ```

* Push branch to github:

   ```
   git push -u origin my-new-branch
   
   ```

* Delete local branch:

   ```
   git branch --delete my-bad-branch
   
   ```

* Reset local to state of remote. Good idea to run with no-act first to see what would be removed:

   ```
   git clean -n -f

   ```
* Keep your branch synced with master:

   ```
   git checkout master
   git pull
   git checkout my-branch
   git merge master

   ```
