
import { BaseResponse, ErrCode } from "./lib/BaseResponse"
import { int64, int32 } from "./lib/less";
import { AuthType } from "./Auth";

/**
 * 创建
 * @method POST
 */
interface Request {

    /**
     * 键值
     */
    key: string

    /**
     * 类型
     */
    type?: AuthType

    /**
     * 值
     */
    value: string

    /**
     * 超时时间(秒)
     */
    expires: int32
}

interface Response extends BaseResponse {
    data?: any
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
