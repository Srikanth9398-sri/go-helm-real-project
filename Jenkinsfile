pipeline {
    agent any

    environment {
        IMAGE_NAME = "go-app"
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/Srikanth9398-sri/go-helm-real-project.git'
            }
        }

        stage('Go Build') {
            steps {
                sh 'go build -o app'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t $IMAGE_NAME:${BUILD_NUMBER} .'
            }
        }

        stage('Helm Deploy') {
            steps {
                sh '''
                  helm upgrade --install go-app ./helm-chart \
                    --set image.repository=$IMAGE_NAME \
                    --set image.tag=${BUILD_NUMBER}
                '''
            }
        }
    }

    post {
        success {
            echo "✅ Deployment successful"
        }
        failure {
            echo "❌ Deployment failed"
        }
    }
}
