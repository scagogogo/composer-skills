import React from 'react'
import { useTranslation } from 'react-i18next'
import { Typography, Tag } from 'antd'
import { AppstoreOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'
import { StaggerGrid, StaggerItem } from '../components/ScrollReveal'

const { Paragraph } = Typography

const showcaseColors = ['#2563EB', '#0891B2', '#16A34A']

const ShowcaseSection: React.FC = () => {
  const { t } = useTranslation()

  const items = t('showcase.items', { returnObjects: true }) as Array<{
    title: string
    description: string
    tags: string[]
    code: string
  }>

  return (
    <section id="showcase" className="section-alt">
      <div style={{ maxWidth: 960, margin: '0 auto' }}>
        <SectionTitle title={t('showcase.title')} subtitle={t('showcase.subtitle')} />

        <StaggerGrid>
          <div style={{ display: 'grid', gap: 20 }}>
            {items.map((item, index) => {
              const color = showcaseColors[index % showcaseColors.length]
              return (
                <StaggerItem key={index}>
                  <div
                    style={{
                      border: '1px solid #E2E8F0',
                      borderRadius: 4,
                      background: '#fff',
                      overflow: 'hidden',
                    }}
                  >
                    {/* Header */}
                    <div
                      style={{
                        display: 'flex',
                        alignItems: 'center',
                        gap: 12,
                        padding: '14px 20px',
                        background: `${color}06`,
                        borderBottom: `1px solid ${color}15`,
                      }}
                    >
                      <div
                        style={{
                          width: 32,
                          height: 32,
                          borderRadius: 4,
                          background: color,
                          display: 'flex',
                          alignItems: 'center',
                          justifyContent: 'center',
                          color: '#fff',
                          fontSize: 16,
                          flexShrink: 0,
                        }}
                      >
                        <AppstoreOutlined />
                      </div>
                      <div style={{ flex: 1 }}>
                        <div style={{ fontWeight: 700, fontSize: 15, color: '#0F172A' }}>
                          {item.title}
                        </div>
                      </div>
                      <div style={{ display: 'flex', gap: 4, flexWrap: 'wrap' }}>
                        {item.tags.map((tag, ti) => (
                          <Tag
                            key={ti}
                            style={{
                              fontSize: 11,
                              borderRadius: 2,
                              margin: 0,
                              padding: '0 6px',
                              background: '#F8FAFC',
                              border: '1px solid #E2E8F0',
                            }}
                          >
                            {tag}
                          </Tag>
                        ))}
                      </div>
                    </div>

                    {/* Body */}
                    <div style={{ padding: '0 20px 20px' }}>
                      <Paragraph
                        style={{ color: '#64748B', margin: '12px 0', lineHeight: 1.6, fontSize: 14 }}
                      >
                        {item.description}
                      </Paragraph>
                      <div className="code-wrapper">
                        <CodeBlock code={item.code} />
                      </div>
                    </div>
                  </div>
                </StaggerItem>
              )
            })}
          </div>
        </StaggerGrid>
      </div>
    </section>
  )
}

export default ShowcaseSection
