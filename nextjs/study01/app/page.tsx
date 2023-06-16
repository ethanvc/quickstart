import Image from 'next/image'

export default function Home() {
    const items1 = ["不限提问次数，只限制总字数", "有效期：7天", "字数：45000个", "支持上下文关联"];
    const items2 = ["不限提问次数，只限制总字数", "有效期：30天", "字数：135000个", "支持上下文关联", "直连模式"];
    return (
        <div>
            <PlanCard items={items1}/>
            <PlanCard items={items2}/>
        </div>
    )
}

function PlanCard({items}: {items:string[]}) {
    return (<div className="h-[280px] w-[280px] border-[1px] border-[rgb(103,194,58)] p-[10px]">
        <div className="h-[calc(100%-32px)]">
            <div className="inline-block text-[rgb(24,160,88)] bg-[rgba(24,160,88,0.12)] p-1">体验套餐</div>
            <div className="text-4xl">¥9.90</div>
            <div className="line-through text-gray-400">原价：19.90</div>
            <ul>
                {items.map((item, index) => (
                    <li><Image
                        src="/tick.svg"
                        alt=""
                        className="inline-block"
                        width={20}
                        height={20}
                    />{item}</li>
                ))}
            </ul>
        </div>
        <button className="h-[32px] block w-[calc(100%-5px)] text-[rgb(24,160,88)] bg-[rgba(24,160,88,0.16)] p-1 hover:bg-[rgba(24,160,88,0.30)]"
                type="button">确认支付
        </button>
    </div>);
}
