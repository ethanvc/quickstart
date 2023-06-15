import Image from 'next/image'

export default function Home() {
    const items = ["不限提问次数", "有效期：7天", "字数：45000个", "支持上下文关联", "直连模式"];
    return (
        <div>
            <div className="w-[280px] border-[1px] border-[rgb(103,194,58)] p-[10px]">
                <div className="inline-block text-[rgb(24,160,88)] bg-[rgba(24,160,88,0.12)] p-1">体验套餐</div>
                <div className="text-4xl">¥9.90</div>
                <div className="line-through text-gray-400">原价：19.90</div>
                <ul>
                    {items.map((item, index) =>(
                        <li><Image
                            src="/tick.svg"
                            alt=""
                            className="inline-block"
                            width={20}
                            height={20}
                        />{item}</li>
                    ))}
                </ul>
                <button className="block">确认支付</button>
            </div>
        </div>
    )
}
