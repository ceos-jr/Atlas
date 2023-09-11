export const Base = new URL("localhost:3000");
export enum FetchErrTypes {
    Rejected,
    NetworkError,
    ParseError
}

export interface ErrFetch {
    type: FetchErrTypes
    msg: string
    source: Error | null
}
const makeErrFetch =
    (type: FetchErrTypes, msg: string, source: Error | null) =>
        ({
            type: type,
            msg: msg,
            source: source
        } as ErrFetch)

export interface Result<T, E> {
    ok: T,
    err: E | null
}

export const fetchMaker = <T, E>(
    url: URL,
    bodyHandler: (data: any) => T,
    catchHandler: (errEF: ErrFetch) => E,
    defaultT: T
) => async (): Promise<Result<T, E>> => {
        let result = {
            ok: defaultT,
            err: null
        } as Result<T, E>

        return fetch(url)
            .then(res => {
                if (res.ok) {
                    return res.json()
                        .then(json => {
                            result.ok = bodyHandler(json);
                            return result;
                        })
                        .catch(err => {
                            result.err = catchHandler(
                                    makeErrFetch(
                                        FetchErrTypes.ParseError,
                                        "Parse to json fail",
                                        err));
                            return result
                            });
                }

                result.err = catchHandler(
                    makeErrFetch(
                        FetchErrTypes.Rejected,
                        "Rejected request",
                        null
                    ))
                return result;
            })
            .catch(err => {
                result.err = catchHandler(
                    makeErrFetch(
                        FetchErrTypes.NetworkError,
                        "Network fail",
                        err));
                return result;
            });
}