pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                sh 'go test -v ./cmd/...'
            }
        }
    }
}

