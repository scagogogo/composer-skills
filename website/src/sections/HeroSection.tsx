import React from 'react'
import { useTranslation } from 'react-i18next'
import { Button, Row, Col, Typography } from 'antd'
import { RocketOutlined, GithubOutlined } from '@ant-design/icons'

const { Title, Paragraph } = Typography

const HeroSection: React.FC = () => {
  const { t } = useTranslation()

  const stats = [
    { value: '234', label: t('hero.statMethods') },
    { value: '20', label: t('hero.statApi') },
    { value: '50+', label: t('hero.statCli') },
    { value: '450+', label: t('hero.statTests') },
  ]

  return (
    <section
      id="hero"
      style={{
        background: 'linear-gradient(135deg, #4F46E5 0%, #0284C7 100%)',
        color: '#fff',
        padding: '100px 24px 80px',
        textAlign: 'center',
      }}
    >
      <div style={{ maxWidth: 900, margin: '0 auto' }}>
        <Title
          level={1}
          style={{
            color: '#fff',
            fontSize: 'clamp(28px, 5vw, 48px)',
            marginBottom: 20,
            lineHeight: 1.2,
          }}
        >
          {t('hero.tagline')}
        </Title>
        <Paragraph
          style={{
            color: 'rgba(255, 255, 255, 0.85)',
            fontSize: 'clamp(16px, 2vw, 20px)',
            maxWidth: 680,
            margin: '0 auto 36px',
            lineHeight: 1.6,
          }}
        >
          {t('hero.subtitle')}
        </Paragraph>

        <div style={{ display: 'flex', gap: 16, justifyContent: 'center', flexWrap: 'wrap' }}>
          <Button
            type="primary"
            size="large"
            icon={<RocketOutlined />}
            href="#quickStart"
            style={{
              background: '#fff',
              color: '#4F46E5',
              border: 'none',
              fontWeight: 600,
              height: 48,
              paddingInline: 32,
              borderRadius: 8,
            }}
          >
            {t('hero.cta')}
          </Button>
          <Button
            size="large"
            icon={<GithubOutlined />}
            href="https://github.com/scagogogo/composer-skills"
            target="_blank"
            style={{
              background: 'transparent',
              color: '#fff',
              border: '1px solid rgba(255, 255, 255, 0.4)',
              fontWeight: 600,
              height: 48,
              paddingInline: 32,
              borderRadius: 8,
            }}
          >
            {t('hero.ctaSecondary')}
          </Button>
        </div>

        <Row
          gutter={[24, 24]}
          style={{ marginTop: 64, maxWidth: 700, marginLeft: 'auto', marginRight: 'auto' }}
        >
          {stats.map((stat, index) => (
            <Col xs={12} sm={6} key={index}>
              <div
                style={{
                  background: 'rgba(255, 255, 255, 0.1)',
                  borderRadius: 12,
                  padding: '20px 12px',
                  backdropFilter: 'blur(10px)',
                }}
              >
                <div style={{ fontSize: 32, fontWeight: 700, lineHeight: 1.2 }}>{stat.value}</div>
                <div style={{ fontSize: 14, color: 'rgba(255, 255, 255, 0.75)', marginTop: 4 }}>
                  {stat.label}
                </div>
              </div>
            </Col>
          ))}
        </Row>
      </div>
    </section>
  )
}

export default HeroSection
