import { AppClient } from './generated'

// TODO: 開発環境と本番環境とでリクエスト先を切り替えられるようにする
const apiClient = new AppClient({
  BASE: 'https://h23s15-prism.trap.show'
})

export default apiClient
