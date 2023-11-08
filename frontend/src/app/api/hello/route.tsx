import {Base, ErrFetch, fetchMaker, Result} from "@/app/api/base";

enum Routes {
    Default
}

interface IHello {
    hello: () => Promise<Result<string, string>>
}
const url = new URL("/hello", Base);

const hello: IHello = {
    hello: fetchMaker(
        url,
        (jsonData: any) => jsonData.body as string,
        (err: ErrFetch) => err.msg,
        "empty"
        )
}
