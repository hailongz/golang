
import { BaseResponse, ErrCode } from "../../../lib/BaseResponse"
import { int64, int32 } from "../../../lib/less";
import { Authority } from '../../../Authority';

/**
 * 角色删除资源
 * @method POST
 */
interface Request {

    /**
     * 角色
     */
    roleId : int64

    /**
     * 资源ID
     */
    resId: int64


}

interface Response extends BaseResponse {
    data?: Authority
}

export function handle(req: Request): Response {
    return {
        errno: ErrCode.OK
    }
}
