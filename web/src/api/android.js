import service from '@/utils/request'

import { ElMessage } from 'element-plus'

const handleFileError = (res, fileName) => {
  if (typeof (res.data) !== 'undefined') {
    if (res.data.type === 'application/json') {
      const reader = new FileReader()
      reader.onload = function() {
        const message = JSON.parse(reader.result).msg
        ElMessage({
          showClose: true,
          message: message,
          type: 'error'
        })
      }
      reader.readAsText(new Blob([res.data]))
    }
  } else {
    var downloadUrl = window.URL.createObjectURL(new Blob([res]))
    var a = document.createElement('a')
    a.style.display = 'none'
    a.href = downloadUrl
    a.download = fileName
    var event = new MouseEvent('click')
    a.dispatchEvent(event)
  }
}

export const runTestcase = (data) => {
  return service({
    url: '/android/runTestcase',
    method: 'post',
    data
  })
}

export const getRuntimeState = (data) => {
  return service({
    url: '/android/getRuntimeState',
    method: 'post',
    data
  })
}

export const jobChanged = (data) => {
  return service({
    url: '/android/jobChanged',
    method: 'post',
    data
  })
}

export const downloadFile = (fileName) => {
  return service({
    url: '/android/downloadFile',
    method: 'get',
    params: {
      fileName: fileName
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, fileName)
  })
}
