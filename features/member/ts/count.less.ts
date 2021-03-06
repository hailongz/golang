
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";

/**
 * 数量
 * @method GET
 */
export interface Request {

    /**
     * 商户ID
     */
    bid?: int64

    /**
     * 成员ID
     */
    uid?: int64

    /**
     * 搜索关键字
     */
    q?: string

}


export interface CountData {
    /**
     * 总记录数
     */
    total: int32
}


export interface Response extends BaseResponse {
    data?: CountData
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
