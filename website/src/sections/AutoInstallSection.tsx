import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Image, Steps } from 'antd'
import { SearchOutlined, CheckCircleOutlined, DownloadOutlined, RocketOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'

const AutoInstallSection: React.FC = () => {
  const { t } = useTranslation()

  const steps = [
    {
      icon: <SearchOutlined />,
      title: t('autoInstall.detectTitle'),
      description: t('autoInstall.detectDesc'),
    },
    {
      icon: <CheckCircleOutlined />,
      title: t('autoInstall.checkTitle'),
      description: t('autoInstall.checkDesc'),
    },
    {
      icon: <DownloadOutlined />,
      title: t('autoInstall.installTitle'),
      description: t('autoInstall.installDesc'),
    },
    {
      icon: <RocketOutlined />,
      title: t('autoInstall.readyTitle'),
      description: t('autoInstall.readyDesc'),
    },
  ]

  return (
    <section id="auto-install">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('autoInstall.title')} subtitle={t('autoInstall.subtitle')} />

        <Row gutter={[32, 32]} align="middle">
          <Col xs={24} lg={14}>
            <div style={{ textAlign: 'center', marginBottom: 24 }}>
              <Image
                src={`${import.meta.env.BASE_URL}images/auto-install-flow.png`}
                alt="Auto-Install Flow"
                style={{ maxWidth: '100%', borderRadius: 8 }}
                preview={false}
              />
            </div>
            <CodeBlock code={t('autoInstall.code')} />
          </Col>
          <Col xs={24} lg={10}>
            <Steps
              direction="vertical"
              current={-1}
              items={steps.map((step, i) => ({
                title: (
                  <span style={{ fontSize: 16, fontWeight: 600 }}>
                    {step.title}
                  </span>
                ),
                description: <span style={{ color: '#475569' }}>{step.description}</span>,
                icon: (
                  <div
                    style={{
                      width: 40,
                      height: 40,
                      borderRadius: '50%',
                      background: `linear-gradient(135deg, ${i % 2 === 0 ? '#4F46E5' : '#0284C7'}, ${i % 2 === 0 ? '#6366F1' : '#0EA5E9'})`,
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      color: '#fff',
                      fontSize: 18,
                    }}
                  >
                    {step.icon}
                  </div>
                ),
              }))}
            />
          </Col>
        </Row>

        <div style={{ textAlign: 'center', marginTop: 48 }}>
          <Image
            src={`${import.meta.env.BASE_URL}images/platform-matrix.png`}
            alt="Platform Matrix"
            style={{ maxWidth: '80%', borderRadius: 8 }}
            preview={false}
          />
        </div>
      </div>
    </section>
  )
}

export default AutoInstallSection
