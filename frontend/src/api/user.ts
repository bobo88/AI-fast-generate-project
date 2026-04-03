import request from '@/utils/request'
import type { User, UserForm, UserQuery, PageResponse } from '@/types/user'

export function getUserList(params: UserQuery) {
  return request.get<PageResponse<User>>('/user/list', { params })
}

export function getUserDetail(id: number) {
  return request.get<User>('/user/detail', { params: { id } })
}

export function createUser(data: UserForm) {
  return request.post('/user/create', data)
}

export function updateUser(data: UserForm & { id: number }) {
  return request.post('/user/update', data)
}

export function deleteUser(id: number) {
  return request.post('/user/delete', { id })
}

export function updatePassword(data: { id: number; old_password: string; new_password: string }) {
  return request.post('/user/update-password', data)
}

export function assignRoles(data: { user_id: number; role_ids: number[] }) {
  return request.post('/user/assign-roles', data)
}
