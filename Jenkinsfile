pipeline {
  agent any
  stages {
    stage('go lint') {
      steps {
        sh ' golint main.go'
      }
    }
  }
}