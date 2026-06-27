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
            <div style={{ textAlign: 'center', marginBottom: 20 }}>
              <Image
                src={`${import.meta.env.BASE_URL}images/auto-install-flow.png`}
                alt="Auto-Install Flow"
                style={{ maxWidth: '100%' }}
                preview={false}
              />
            </div>
            <CodeBlock code={t('autoInstall.code')} />
          </Col>
          <Col xs={24} lg={10}>
            <Steps
              direction="vertical"
              current={-1}
              items={steps.map((step) => ({
                title: (
                  <span style={{ fontSize: 15, fontWeight: 600 }}>
                    {step.title}
                  </span>
                ),
                description: <span style={{ color: '#64748B' }}>{step.description}</span>,
                icon: (
                  <div
                    style={{
                      width: 32,
                      height: 32,
                      borderRadius: 4,
                      background: '#2563EB',
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      color: '#fff',
                      fontSize: 15,
                    }}
                  >
                    {step.icon}
                  </div>
                ),
              }))}
            />
          </Col>
        </Row>

        <div style={{ textAlign: 'center', marginTop: 40 }}>
          <Image
            src={`${import.meta.env.BASE_URL}images/platform-matrix.png`}
            alt="Platform Matrix"
            style={{ maxWidth: '80%' }}
            preview={false}
          />
        </div>
      </div>
    </section>
  )
}

export default AutoInstallSection
