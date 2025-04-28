#!/usr/bin/env groovy

pipeline {
    //-------
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building...'
                sh 'docker stop simplehttp'
                sh 'docker rm simplehttp'
                sh 'docker rmi simplehttp'
                sh 'docker build -t simplehttp .'
                sh 'docker run -dit --name simplehttp -p 8082:8082 simplehttp'
            }
        }
    }
}
