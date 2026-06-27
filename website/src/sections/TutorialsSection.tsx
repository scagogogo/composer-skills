import React from 'react'
import { useTranslation } from 'react-i18next'
import { Tag, Typography } from 'antd'
import { ClockCircleOutlined, ReadOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import { StaggerGrid, StaggerItem } from '../components/ScrollReveal'

const { Paragraph } = Typography

const difficultyConfig: Record<string, { color: string; level: number }> = {
  'Beginner': { color: '#16A34A', level: 1 },
  '入门': { color: '#16A34A', level: 1 },
  'Intermediate': { color: '#2563EB', level: 2 },
  '进阶': { color: '#2563EB', level: 2 },
  'Advanced': { color: '#DC2626', level: 3 },
  '高级': { color: '#DC2626', level: 3 },
}

const TutorialsSection: React.FC = () => {
  const { t } = useTranslation()

  const items = t('tutorials.items', { returnObjects: true }) as Array<{
    category: string
    categoryColor: string
    title: string
    description: string
    difficulty: string
    readTime: string
  }>

  return (
    <section id="tutorials">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('tutorials.title')} subtitle={t('tutorials.subtitle')} />

        <StaggerGrid>
          <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(320px, 1fr))', gap: 20 }}>
            {items.map((item, index) => {
              const diff = difficultyConfig[item.difficulty] || { color: '#475569', level: 1 }
              return (
                <StaggerItem key={index}>
                  <div
                    className="tutorial-card"
                    style={{
                      border: '1px solid #E2E8F0',
                      borderRadius: 4,
                      padding: 24,
                      background: '#fff',
                      transition: 'border-color 0.2s ease, transform 0.2s ease',
                    }}
                    onMouseEnter={(e) => {
                      e.currentTarget.style.borderColor = item.categoryColor
                      e.currentTarget.style.transform = 'translateY(-2px)'
                    }}
                    onMouseLeave={(e) => {
                      e.currentTarget.style.borderColor = '#E2E8F0'
                      e.currentTarget.style.transform = 'translateY(0)'
                    }}
                  >
                    {/* Top row: category + difficulty */}
                    <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 16 }}>
                      <Tag
                        style={{
                          background: `${item.categoryColor}0A`,
                          color: item.categoryColor,
                          border: `1px solid ${item.categoryColor}30`,
                          fontWeight: 600,
                          borderRadius: 2,
                          margin: 0,
                          fontSize: 12,
                        }}
                      >
                        {item.category}
                      </Tag>
                      <div style={{ display: 'flex', alignItems: 'center', gap: 4 }}>
                        {Array.from({ length: 3 }).map((_, i) => (
                          <div
                            key={i}
                            style={{
                              width: 6,
                              height: 6,
                              borderRadius: 1,
                              background: i < diff.level ? diff.color : '#E2E8F0',
                            }}
                          />
                        ))}
                      </div>
                    </div>

                    {/* Title */}
                    <div style={{ fontWeight: 700, fontSize: 15, color: '#0F172A', marginBottom: 8, lineHeight: 1.4 }}>
                      <ReadOutlined style={{ marginRight: 6, color: '#2563EB' }} />
                      {item.title}
                    </div>

                    {/* Description */}
                    <Paragraph style={{ color: '#64748B', marginBottom: 14, lineHeight: 1.6, fontSize: 14 }}>
                      {item.description}
                    </Paragraph>

                    {/* Bottom: read time + difficulty label */}
                    <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                      <div style={{ display: 'flex', alignItems: 'center', color: '#94A3B8', fontSize: 12, gap: 4 }}>
                        <ClockCircleOutlined />
                        {item.readTime}
                      </div>
                      <Tag
                        style={{
                          background: '#F8FAFC',
                          color: diff.color,
                          border: `1px solid ${diff.color}30`,
                          borderRadius: 2,
                          margin: 0,
                          fontWeight: 500,
                          fontSize: 12,
                        }}
                      >
                        {item.difficulty}
                      </Tag>
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

export default TutorialsSection
