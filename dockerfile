# Docker Build
- name: 'gcr.io/cloud-builders/docker'
  args: ['build', '-t', 
         'us-central1-docker.pkg.dev/$PROJECT_ID/$_REPO_NAME/myimage:$SHORT_SHA', '.']
      
# Docker push to Google Artifact Registry
- name: 'gcr.io/cloud-builders/docker'
  args: ['push', 'us-central1-docker.pkg.dev/$PROJECT_ID/$_REPO_NAME/myimage:$SHORT_SHA']
  
# Deploy to Cloud Run
- name: 'gcr.io/cloud-builders/gcloud'
  args: ['run', 'deploy', 'helloworld', 
         '--image=us-central1-docker.pkg.dev/$PROJECT_ID/$_REPO_NAME/myimage:$SHORT_SHA', 
         '--region', 'us-central1', '--platform', 'managed']
