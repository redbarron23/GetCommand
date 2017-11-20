pipeline {
  agent any
  stages {
    stage('go lint') {
      steps {
        sh ' /Users/jhourihane/work/bin/golint main.go'
      }
    }
    stage('go vet') {
      steps {
        sh '/usr/local/bin/go vet'
      }
    }
    stage('go build') {
      steps {
        sh '/usr/local/bin/go build'
      }
    }
    stage('run executable') {
      steps {
        sh 'GetCommand'
      }
    }
  }
}