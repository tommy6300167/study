function checkStatus() {
    const statusElement = document.getElementById('status');
    statusElement.textContent = '正在檢查...';
    statusElement.className = '';
    
    fetch('/api/status')
        .then(response => {
            if (response.ok) {
                statusElement.textContent = '✅ 伺服器運行正常';
                statusElement.className = 'success';
            } else {
                statusElement.textContent = '❌ 伺服器回應異常';
                statusElement.className = 'error';
            }
        })
        .catch(error => {
            statusElement.textContent = '❌ 無法連接到伺服器';
            statusElement.className = 'error';
        });
}

function sendMessage() {
    const responseElement = document.getElementById('response');
    responseElement.textContent = '發送中...';
    
    fetch('/api/hello', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({message: 'Hello from frontend!'})
    })
    .then(response => response.text())
    .then(data => {
        responseElement.textContent = `伺服器回應: ${data}`;
        responseElement.className = 'success';
    })
    .catch(error => {
        responseElement.textContent = `錯誤: ${error.message}`;
        responseElement.className = 'error';
    });
}

// 頁面載入時自動檢查狀態
document.addEventListener('DOMContentLoaded', checkStatus);