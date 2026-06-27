import React from 'react'
import { useTranslation } from 'react-i18next'
import { Row, Col, Typography } from 'antd'
import { CloudOutlined, CodeOutlined, CheckCircleOutlined, CloseCircleOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import { RevealSection, StaggerGrid, StaggerItem } from '../components/ScrollReveal'

const { Paragraph } = Typography

const SDKComparisonSection: React.FC = () => {
  const { t } = useTranslation()

  const fields = t('sdkComparison.fields', { returnObjects: true }) as Array<{
    label: string
    packagist: string
    composer: string
  }>

  return (
    <section id="sdk-comparison">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('sdkComparison.title')} subtitle={t('sdkComparison.subtitle')} />

        {/* Two SDK cards side by side */}
        <RevealSection>
          <Row gutter={[20, 20]} style={{ marginBottom: 32 }}>
            <Col xs={24} md={12}>
              <div
                style={{
                  border: '1px solid #A5F3FC',
                  background: '#ECFEFF',
                  borderRadius: 4,
                  padding: 24,
                  height: '100%',
                }}
              >
                <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 16 }}>
                  <div
                    style={{
                      width: 36,
                      height: 36,
                      borderRadius: 4,
                      background: '#0891B2',
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      color: '#fff',
                      fontSize: 18,
                    }}
                  >
                    <CloudOutlined />
                  </div>
                  <div style={{ fontWeight: 700, fontSize: 16, color: '#0E7490' }}>
                    {t('sdkComparison.packagistTitle')}
                  </div>
                </div>
                <Paragraph style={{ color: '#475569', margin: 0, fontSize: 14, lineHeight: 1.6 }}>
                  HTTP calls to Packagist API. Pure Go — no PHP required. Search packages, get stats, security advisories.
                </Paragraph>
              </div>
            </Col>
            <Col xs={24} md={12}>
              <div
                style={{
                  border: '1px solid #BFDBFE',
                  background: '#EFF6FF',
                  borderRadius: 4,
                  padding: 24,
                  height: '100%',
                }}
              >
                <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 16 }}>
                  <div
                    style={{
                      width: 36,
                      height: 36,
                      borderRadius: 4,
                      background: '#2563EB',
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                      color: '#fff',
                      fontSize: 18,
                    }}
                  >
                    <CodeOutlined />
                  </div>
                  <div style={{ fontWeight: 700, fontSize: 16, color: '#1D4ED8' }}>
                    {t('sdkComparison.composerTitle')}
                  </div>
                </div>
                <Paragraph style={{ color: '#475569', margin: 0, fontSize: 14, lineHeight: 1.6 }}>
                  Executes local composer binary. Requires PHP 7.4+ and Composer 2.0+. Install/update deps, manage projects, audit, run scripts.
                </Paragraph>
              </div>
            </Col>
          </Row>
        </RevealSection>

        {/* Comparison rows */}
        <StaggerGrid>
          <div style={{ border: '1px solid #E2E8F0', borderRadius: 4, overflow: 'hidden' }}>
            {/* Header row */}
            <div
              style={{
                display: 'grid',
                gridTemplateColumns: '1fr 1fr 1fr',
                background: '#F8FAFC',
                borderBottom: '2px solid #E2E8F0',
                padding: '10px 16px',
              }}
            >
              <div style={{ fontWeight: 600, fontSize: 13, color: '#475569' }}>Field</div>
              <div style={{ fontWeight: 600, fontSize: 13, color: '#0891B2', textAlign: 'center' }}>
                <CloudOutlined style={{ marginRight: 4 }} />Packagist
              </div>
              <div style={{ fontWeight: 600, fontSize: 13, color: '#2563EB', textAlign: 'center' }}>
                <CodeOutlined style={{ marginRight: 4 }} />Composer
              </div>
            </div>
            {/* Data rows */}
            {fields.map((field, index) => (
              <StaggerItem key={index}>
                <div
                  style={{
                    display: 'grid',
                    gridTemplateColumns: '1fr 1fr 1fr',
                    padding: '12px 16px',
                    borderBottom: index < fields.length - 1 ? '1px solid #F1F5F9' : 'none',
                    background: '#fff',
                  }}
                >
                  <div style={{ fontWeight: 600, fontSize: 13, color: '#0F172A' }}>{field.label}</div>
                  <div style={{ fontSize: 13, color: '#475569', textAlign: 'center' }}>
                    {field.label === 'Requires PHP?' || field.label === '需要 PHP？' ? (
                      field.packagist.includes('No') || field.packagist.includes('不需要') ? (
                        <span style={{ color: '#16A34A' }}><CheckCircleOutlined style={{ marginRight: 4 }} />{field.packagist}</span>
                      ) : (
                        field.packagist
                      )
                    ) : (
                      field.packagist
                    )}
                  </div>
                  <div style={{ fontSize: 13, color: '#475569', textAlign: 'center' }}>
                    {field.label === 'Requires PHP?' || field.label === '需要 PHP？' ? (
                      field.composer.includes('Yes') || field.composer.includes('需要') ? (
                        <span style={{ color: '#EA580C' }}><CloseCircleOutlined style={{ marginRight: 4 }} />{field.composer}</span>
                      ) : (
                        field.composer
                      )
                    ) : (
                      field.composer
                    )}
                  </div>
                </div>
              </StaggerItem>
            ))}
          </div>
        </StaggerGrid>
      </div>
    </section>
  )
}

export default SDKComparisonSection
