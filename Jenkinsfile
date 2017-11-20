pipeline {
  agent any
  stages {
    stage('go lint') {
      parallel {
        stage('git pull') {
          steps {
            sh 'git pull'
          }
        }
        stage('golint') {
          steps {
            sh '/Users/jhourihane/work/bin/golint main.go'
          }
        }
      }
    }
    stage('go vet') {
      steps {
        sh '/usr/local/bin/go vet'
      }
    }
    stage('go build') {
      steps {
        sh '/usr/local/bin/go build -o GetCommand'
      }
    }
    stage('run executable') {
      parallel {
        stage('run executable') {
          steps {
            sh '''

echo $WORKSPACE'''
          }
        }
        stage('error') {
          steps {
            sh '''
$WORKSPACE/GetCommand'''
          }
        }
      }
    }
    stage('cleanup') {
      steps {
        sh 'rm $WORKSPACE/GetCommand'
      }
    }
  }
}