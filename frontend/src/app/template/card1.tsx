import { Card1Type } from "./page"

type Card1Props = Card1Type

//+1 prop: Assignees (como?)
const Card1 = ({ taskName, sectorName } : Card1Props) => {
    return (
        <div className="flex overflow-hidden flex-col items-center bg-gray-100 rounded-lg shadow-xl min-h-64"> 
            <div className="p-6">
                <h1 className="text-2xl font-bold text-gray-800">{taskName}</h1>
                <p className="text-gray-500">{sectorName}</p>
            </div>
        </div>
    )
}
export default Card1

//ServiceCard=Card1
