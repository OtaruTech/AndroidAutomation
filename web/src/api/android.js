import service from '@/utils/request'

export const runTestcase = (data) => {
  return service({
    url: '/android/runTestcase',
    method: 'post',
    data
  })
}

export const getRuntimeState = (data) => {
  console.log(data)
  return service({
    url: '/android/getRuntimeState',
    method: 'post',
    data
  })
}
