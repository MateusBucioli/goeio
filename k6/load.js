import http from 'k6/http'
import { check, sleep } from 'k6'

export const options = {
    stages: [
        { duration: '5s', target: 200 },
        { duration: '30s', target: 1000 },
        { duration: '5s', target: 0 },
    ],
}

export default function () {
    const BASE_URL = 'http://localhost:8080'

    const payload = JSON.stringify({
        message: 'Potato!',
    })

    const headers = {
        'Content-Type': 'application/json',
    }

    const response = http.post(`${BASE_URL}/send`, payload, { headers })

    check(response, { 'status is 200': (r) => r.status === 200 })

    sleep(1)
}
