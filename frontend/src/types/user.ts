export interface User {
  id: number
  username: string
  nickname?: string
  email?: string
  phone?: string
  avatar?: string
  status: number
  created_at: string
  updated_at: string
  roles?: Role[]
}

export interface Role {
  id: number
  name: string
  code: string
}

export interface UserForm {
  username: string
  password?: string
  nickname?: string
  email?: string
  phone?: string
  status?: number
}

export interface UserQuery {
  username?: string
  nickname?: string
  email?: string
  phone?: string
  status?: number
  page?: number
  page_size?: number
}

export interface PageResponse<T> {
  total: number
  page: number
  page_size: number
  list: T[]
}
