import React from 'react'
import { useTranslation } from 'react-i18next'
import { Typography, Collapse, Tag } from 'antd'
import { AppstoreOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'

const { Title, Paragraph } = Typography

const ShowcaseSection: React.FC = () => {
  const { t } = useTranslation()

  const items = t('showcase.items', { returnObjects: true }) as Array<{
    title: string
    description: string
    tags: string[]
    code: string
  }>

  const collapseItems = items.map((item, index) => ({
    key: String(index),
    label: (
      <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
        <div
          style={{
            width: 32,
            height: 32,
            borderRadius: 4,
            background: '#EFF6FF',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            color: '#2563EB',
            fontSize: 16,
          }}
        >
          <AppstoreOutlined />
        </div>
        <div>
          <Title level={5} style={{ margin: 0, fontSize: 14 }}>{item.title}</Title>
          <div style={{ display: 'flex', gap: 4, marginTop: 4 }}>
            {item.tags.map((tag, ti) => (
              <Tag key={ti} style={{ fontSize: 11, borderRadius: 2, margin: 0, padding: '0 4px' }}>
                {tag}
              </Tag>
            ))}
          </div>
        </div>
      </div>
    ),
    children: (
      <div>
        <Paragraph style={{ color: '#64748B', marginBottom: 12, lineHeight: 1.6, fontSize: 14 }}>
          {item.description}
        </Paragraph>
        <div className="code-wrapper">
          <CodeBlock code={item.code} />
        </div>
      </div>
    ),
  }))

  return (
    <section id="showcase" className="section-alt">
      <div style={{ maxWidth: 900, margin: '0 auto' }}>
        <SectionTitle title={t('showcase.title')} subtitle={t('showcase.subtitle')} />

        <Collapse
          defaultActiveKey={['0']}
          items={collapseItems}
          style={{ background: '#fff' }}
        />
      </div>
    </section>
  )
}

export default ShowcaseSection
