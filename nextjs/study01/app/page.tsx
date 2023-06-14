import Image from 'next/image'

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="w-[280px] border-[1px] border-[rgb(103,194,58)] p-[10px]">
        <div>哈弗大学</div>
          <div>常规通道</div>
          <div>走特殊通道</div>
          <div>
            <ul>
                <li>有没的校园环境</li>
                <li>资深的教学团队</li>
                <li>额外的水果套餐</li>
                <li>有专门的人辅导你</li>
            </ul>
          </div>
      </div>
    </main>
  )
}
