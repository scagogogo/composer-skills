import React from 'react'
import { Card, Typography } from 'antd'

const { Title, Paragraph } = Typography

interface FeatureCardProps {
  icon: React.ReactNode
  title: string
  description: string
}

const FeatureCard: React.FC<FeatureCardProps> = ({ icon, title, description }) => {
  return (
    <Card
      hoverable
      style={{
        height: '100%',
        borderRadius: 12,
        border: '1px solid #E2E8F0',
      }}
      styles={{
        body: {
          padding: '24px',
        },
      }}
    >
      <div style={{ fontSize: 32, color: '#4F46E5', marginBottom: 16 }}>{icon}</div>
      <Title level={5} style={{ marginBottom: 8 }}>
        {title}
      </Title>
      <Paragraph style={{ color: '#475569', marginBottom: 0 }}>{description}</Paragraph>
    </Card>
  )
}

export default FeatureCard
