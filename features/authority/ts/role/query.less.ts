
import { BaseResponse, ErrCode } from "../lib/BaseResponse"
import { int64, int32 } from "../lib/less";
import { Page } from '../Query';
import { Role } from "../Role";

/**
 * 查询角色
 * @method GET
 */
export interface Request {

    /**
     * 前缀
     */
    prefix?: string

    /**
     * 分页位置, 从1开始, 0 不处理分页
     */
    p?: int32

    /**
     * 分页大小，默认 20
     */
    n?: int32

}

export interface RoleQueryData {

    /**
     * 资源
     */
    items: Role[]

    /**
     * 分页
     */
    page?: Page
}


export interface Response extends BaseResponse {
    data?: RoleQueryData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
