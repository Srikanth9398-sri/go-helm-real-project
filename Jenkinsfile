pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                echo 'Code checkout successful'
            }
        }

        stage('Deploy with Helm') {
            steps {
                sh '''
                helm upgrade --install go-helm-app ./helm-chart
                '''
            }
        }
    }
}
