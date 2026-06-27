import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Tag, Typography } from 'antd'
import {
  ClockCircleOutlined,
  ReadOutlined,
} from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'

const { Title, Paragraph } = Typography

const difficultyColors: Record<string, string> = {
  'Beginner': 'green',
  '入门': 'green',
  'Intermediate': 'blue',
  '进阶': 'blue',
  'Advanced': 'red',
  '高级': 'red',
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

        <Row gutter={[24, 24]}>
          {items.map((item, index) => (
            <Col xs={24} sm={12} lg={8} key={index}>
              <div className="tutorial-card">
                <div style={{ display: 'flex', alignItems: 'center', justifyContent: 'space-between', marginBottom: 16 }}>
                  <Tag
                    style={{
                      background: `${item.categoryColor}10`,
                      color: item.categoryColor,
                      border: `1px solid ${item.categoryColor}20`,
                      fontWeight: 600,
                      borderRadius: 6,
                      margin: 0,
                    }}
                  >
                    {item.category}
                  </Tag>
                  <Tag
                    color={difficultyColors[item.difficulty] || 'default'}
                    style={{ borderRadius: 6, margin: 0, fontWeight: 500 }}
                  >
                    {item.difficulty}
                  </Tag>
                </div>

                <Title level={5} style={{ marginBottom: 10, fontWeight: 700, fontSize: 16, lineHeight: 1.4 }}>
                  <ReadOutlined style={{ marginRight: 8, color: '#4F46E5' }} />
                  {item.title}
                </Title>

                <Paragraph style={{ color: '#64748B', marginBottom: 16, lineHeight: 1.6, fontSize: 14 }}>
                  {item.description}
                </Paragraph>

                <div style={{ display: 'flex', alignItems: 'center', color: '#94A3B8', fontSize: 13 }}>
                  <ClockCircleOutlined style={{ marginRight: 6 }} />
                  {item.readTime}
                </div>
              </div>
            </Col>
          ))}
        </Row>
      </div>
    </section>
  )
}

export default TutorialsSection
