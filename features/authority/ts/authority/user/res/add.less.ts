
import { BaseResponse, ErrCode } from "../../../lib/BaseResponse"
import { int64, int32 } from "../../../lib/less";
import { Authority } from '../../../Authority';

/**
 * 用户添加资源
 * @method POST
 */
interface Request {

    /**
     * 用户ID
     */
    uid : int64

    /**
     * 资源ID
     */
    resId: int64

    /**
     * 其他选项 JSON 叠加
     */
    options?: string
}

interface Response extends BaseResponse {
    data?: Authority
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
