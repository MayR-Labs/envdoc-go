# envdoc Landing Page Specifications

This document provides comprehensive specifications for building a landing page for the envdoc CLI tool. It includes content, structure, branding guidelines, and technical details to help designers and developers create an effective landing page.

---

## Table of Contents

1. [Project Overview](#project-overview)
2. [Brand Identity](#brand-identity)
3. [Page Structure](#page-structure)
4. [Hero Section](#hero-section)
5. [Features Section](#features-section)
6. [Use Cases Section](#use-cases-section)
7. [Installation Section](#installation-section)
8. [Commands Reference](#commands-reference)
9. [Code Examples](#code-examples)
10. [Testimonials/Social Proof](#testimonialssocial-proof)
11. [Call to Action](#call-to-action)
12. [Footer](#footer)
13. [Technical Specifications](#technical-specifications)
14. [SEO Metadata](#seo-metadata)
15. [Assets Needed](#assets-needed)

---

## Project Overview

**Product Name:** envdoc  
**Tagline:** "Tame Your Environment Variables"  
**Short Description:** A powerful CLI tool for managing, validating, and transforming environment variable files.

**Long Description:**
envdoc is a command-line tool that solves the everyday challenges developers face with environment variables. From generating documentation and schemas to encrypting secrets and keeping environments in sync, envdoc handles the tedious work so you can focus on building.

**Target Audience:**
- Software Developers
- DevOps Engineers
- Security Teams
- System Administrators
- Technical Leads managing multi-environment deployments

**Repository:** https://github.com/MayR-Labs/envdoc-go  
**License:** MIT  
**Version:** 0.1.0

---

## Brand Identity

### Brand Voice
- **Professional** but approachable
- **Technical** but not intimidating
- **Problem-focused** - addresses real pain points
- **Confident** - the tool is well-built and reliable

### Color Palette (Suggested)
| Color Name       | Hex Code  | Usage                          |
|------------------|-----------|--------------------------------|
| Primary Blue     | `#2563EB` | Primary actions, links         |
| Primary Dark     | `#1E40AF` | Hover states, accents          |
| Success Green    | `#10B981` | Success states, checkmarks     |
| Warning Yellow   | `#F59E0B` | Warnings, highlights           |
| Error Red        | `#EF4444` | Errors, critical info          |
| Dark Gray        | `#1F2937` | Text, dark backgrounds         |
| Light Gray       | `#F3F4F6` | Backgrounds, borders           |
| White            | `#FFFFFF` | Backgrounds, text on dark      |

### Typography
- **Headings:** Inter, Poppins, or system sans-serif (bold weight)
- **Body:** Inter, Open Sans, or system sans-serif (regular weight)
- **Code:** JetBrains Mono, Fira Code, or system monospace

### Logo
- Text-based logo using "envdoc" in lowercase
- Consider using a terminal/CLI-inspired design element
- Color: Primary Blue or Dark Gray

---

## Page Structure

### Recommended Page Sections (Top to Bottom)
1. **Navigation Bar** - Logo, nav links, CTA button
2. **Hero Section** - Main value proposition
3. **Problem Statement** - Pain points this tool solves
4. **Features Grid** - Key features with icons
5. **Installation** - Quick start instructions
6. **Commands Showcase** - Interactive command examples
7. **Use Cases** - Real-world scenarios
8. **Code Examples** - Before/after demonstrations
9. **Social Proof** - Stars count, testimonials
10. **Call to Action** - Final push to install
11. **Footer** - Links, credits, legal

---

## Hero Section

### Headline Options
1. "Tame Your Environment Variables"
2. "Stop Fighting with .env Files"
3. "Environment Variable Management, Simplified"
4. "One Tool. All Your .env Problems. Solved."

### Subheadline
"A powerful CLI tool for managing, validating, and transforming environment variable files. Document, sync, encrypt, and convert with ease."

### Primary CTA
- **Text:** "Get Started" or "Install Now"
- **Link:** Scroll to installation section

### Secondary CTA
- **Text:** "View on GitHub"
- **Link:** https://github.com/MayR-Labs/envdoc-go

### Hero Visual
- Terminal window showing envdoc in action
- Animated typing effect (optional)
- Example commands: `envdoc create-example .env` → output

---

## Features Section

### Section Title
"Everything You Need to Manage Environment Files"

### Feature Cards (9 Total)

#### 1. Documentation Generation
- **Icon:** 📝 (Document/Pencil)
- **Title:** "Documentation Generation"
- **Description:** "Create example files and JSON schemas from your .env files automatically. Never leave new team members guessing."
- **Command:** `envdoc create-example .env`

#### 2. File Auditing
- **Icon:** 🔍 (Magnifying Glass)
- **Title:** "File Auditing"
- **Description:** "Find duplicate keys and missing variables across multiple files. Catch configuration errors before they cause problems."
- **Command:** `envdoc audit .env`

#### 3. Environment Sync
- **Icon:** 🔄 (Sync Arrows)
- **Title:** "Environment Synchronization"
- **Description:** "Keep environment files in sync across development, staging, and production. No more environment drift."
- **Command:** `envdoc sync .env .env.staging .env.production`

#### 4. Security & Encryption
- **Icon:** 🔐 (Lock)
- **Title:** "Security & Encryption"
- **Description:** "Encrypt files with AES-256 and generate SHA256 hashes. Store production secrets safely in your repository."
- **Command:** `envdoc encrypt .env.production`

#### 5. Format Conversion
- **Icon:** 🔀 (Shuffle/Convert)
- **Title:** "Format Conversion"
- **Description:** "Convert between .env, JSON, and YAML formats instantly. No more writing custom conversion scripts."
- **Command:** `envdoc to json .env`

#### 6. Schema Validation
- **Icon:** ✅ (Checkmark)
- **Title:** "Schema Validation"
- **Description:** "Validate .env files against JSON schemas. Perfect for CI/CD pipelines and automated deployments."
- **Command:** `envdoc validate .env schema.json`

#### 7. Interactive Experience
- **Icon:** 🎨 (Palette/Interactive)
- **Title:** "Interactive Experience"
- **Description:** "User-friendly prompts and PIN-based confirmations for destructive operations. Safe by design."
- **Command:** Interactive prompts built-in

#### 8. Comprehensive Reports
- **Icon:** 📊 (Chart)
- **Title:** "Comprehensive Reports"
- **Description:** "Generate detailed markdown reports with table of contents. Share findings with your team easily."
- **Command:** `envdoc doctor`

#### 9. Cross-Platform
- **Icon:** 🎯 (Target)
- **Title:** "Cross-Platform"
- **Description:** "Works on Linux, macOS, and Windows. AMD64 and ARM64 architectures. One binary, no dependencies."
- **Platforms:** Linux, macOS, Windows

---

## Use Cases Section

### Section Title
"Built for Real-World Development Teams"

### Use Case 1: For Developers
**Title:** "For Developers"
**Icon:** 💻
**Benefits:**
- Quickly create .env.example files for new team members
- Validate local environment against production schema
- Keep track of required environment variables
- Convert configurations between formats

### Use Case 2: For DevOps
**Title:** "For DevOps"
**Icon:** ⚙️
**Benefits:**
- Audit environment configurations across multiple deployments
- Ensure consistency between staging and production
- Generate documentation for environment variables
- Automate configuration validation in CI/CD

### Use Case 3: For Security Teams
**Title:** "For Security Teams"
**Icon:** 🛡️
**Benefits:**
- Encrypt sensitive configuration files
- Verify file integrity with hash generation
- Track changes in environment configurations
- Audit for accidentally committed secrets

---

## Installation Section

### Section Title
"Get Started in Seconds"

### Quick Install (Linux/macOS)

```bash
# Using curl
curl -sSL https://raw.githubusercontent.com/MayR-Labs/envdoc-go/main/install.sh | bash

# Or using wget
wget -qO- https://raw.githubusercontent.com/MayR-Labs/envdoc-go/main/install.sh | bash
```

### Manual Installation

#### Linux
```bash
# AMD64
wget https://github.com/MayR-Labs/envdoc-go/releases/latest/download/envdoc-linux-amd64
chmod +x envdoc-linux-amd64
sudo mv envdoc-linux-amd64 /usr/local/bin/envdoc

# ARM64
wget https://github.com/MayR-Labs/envdoc-go/releases/latest/download/envdoc-linux-arm64
chmod +x envdoc-linux-arm64
sudo mv envdoc-linux-arm64 /usr/local/bin/envdoc
```

#### macOS
```bash
# Intel
wget https://github.com/MayR-Labs/envdoc-go/releases/latest/download/envdoc-darwin-amd64
chmod +x envdoc-darwin-amd64
sudo mv envdoc-darwin-amd64 /usr/local/bin/envdoc

# Apple Silicon (M1/M2/M3)
wget https://github.com/MayR-Labs/envdoc-go/releases/latest/download/envdoc-darwin-arm64
chmod +x envdoc-darwin-arm64
sudo mv envdoc-darwin-arm64 /usr/local/bin/envdoc
```

#### Windows
Download the latest `envdoc-windows-amd64.exe` from the [releases page](https://github.com/MayR-Labs/envdoc-go/releases/latest) and add it to your PATH.

### Build from Source
```bash
git clone https://github.com/MayR-Labs/envdoc-go.git
cd envdoc-go
go build -o envdoc .
sudo mv envdoc /usr/local/bin/
```

### Verify Installation
```bash
envdoc --version
```

---

## Commands Reference

### Complete Command List

| Command | Description | Example |
|---------|-------------|---------|
| `create-example` | Generate example file with empty values | `envdoc create-example .env` |
| `create-schema` | Generate JSON schema from .env file | `envdoc create-schema .env` |
| `arrange` | Sort and group environment variables | `envdoc arrange .env` |
| `audit` | Find duplicate keys and missing values | `envdoc audit .env` |
| `compare` | Compare keys across multiple files | `envdoc compare .env.dev .env.prod` |
| `sync` | Synchronize keys across multiple files | `envdoc sync .env.dev .env.prod` |
| `clear-values` | Clear all values (keeps keys) | `envdoc clear-values .env` |
| `encrypt` | Encrypt file with AES-256 | `envdoc encrypt .env.production` |
| `decrypt` | Decrypt encrypted file | `envdoc decrypt .env.encrypted` |
| `hash` | Generate SHA256 hash | `envdoc hash .env` |
| `base64` | Encode/decode base64 | `envdoc base64 encode .env` |
| `to` | Convert to JSON/YAML | `envdoc to json .env` |
| `from` | Convert from JSON/YAML to .env | `envdoc from config.json` |
| `validate` | Validate against JSON schema | `envdoc validate .env schema.json` |
| `doctor` | Audit all .env files in directory | `envdoc doctor` |
| `engineer` | Sync and arrange all .env files | `envdoc engineer` |
| `version` | Show version | `envdoc version` |
| `documentation` | Open documentation | `envdoc documentation` |
| `license` | Show license | `envdoc license` |
| `changelog` | Show changelog | `envdoc changelog` |
| `authors` | Show authors | `envdoc authors` |

---

## Code Examples

### Example 1: Project Setup

**Scenario:** Setting up environment configuration for a new project

**Before:**
```
# Manual process:
# 1. Copy .env file (might contain secrets!)
# 2. Manually clear values
# 3. Hope you didn't miss anything
# 4. New developer still asks questions
```

**After:**
```bash
# Create a safe template for new developers
envdoc create-example .env.production .env.example

# Create a schema for validation
envdoc create-schema .env.production .env.schema.json

# Validate any environment file
envdoc validate .env.staging .env.schema.json
```

**Output (.env.example):**
```env
DATABASE_HOST=
DATABASE_PORT=
DATABASE_NAME=
DATABASE_USER=
DATABASE_PASSWORD=

API_KEY=
API_SECRET=
API_BASE_URL=
```

### Example 2: Multi-Environment Management

**Scenario:** Managing multiple environment files

```bash
# Compare environments to see differences
envdoc compare .env.development .env.staging .env.production

# Sync missing keys across all files
envdoc sync .env.development .env.staging .env.production

# Arrange all files for consistency
envdoc arrange .env.development
envdoc arrange .env.staging
envdoc arrange .env.production
```

### Example 3: Security Workflow

**Scenario:** Securing production secrets

```bash
# Encrypt production secrets before committing
envdoc encrypt .env.production
# Outputs: .env.production.encrypted

# Generate hash for verification
envdoc hash .env.production.encrypted
# Outputs: SHA256 hash

# Later, decrypt when needed
envdoc decrypt .env.production.encrypted
```

### Example 4: CI/CD Integration

**Scenario:** Automated validation in pipelines

```yaml
# GitHub Actions Example
steps:
  - name: Validate Environment
    run: |
      envdoc validate .env $SCHEMA_FILE || exit 1
      envdoc compare .env .env.example || exit 1
```

---

## Testimonials/Social Proof

### Stats to Display
- GitHub Stars count (dynamic badge)
- Number of releases
- Supported platforms count (3)
- Supported architectures count (2)

### GitHub Badges
```markdown
[![Release](https://img.shields.io/github/v/release/MayR-Labs/envdoc-go)](https://github.com/MayR-Labs/envdoc-go/releases)
[![License](https://img.shields.io/github/license/MayR-Labs/envdoc-go)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/MayR-Labs/envdoc-go)](go.mod)
```

### Value Propositions
- "Save hours of onboarding time"
- "Catch configuration errors before production"
- "Keep environments in sync automatically"
- "Secure your secrets with AES-256 encryption"

---

## Call to Action

### Primary CTA Section

**Headline:** "Ready to Stop Wasting Time on Environment Variable Chaos?"

**Subheadline:** "Install envdoc in seconds and start managing your .env files like a pro."

**CTA Buttons:**
1. **Primary:** "Install Now" → Scroll to installation
2. **Secondary:** "View on GitHub" → Repository link

### Final Value Reminder
- ✓ Free and open source (MIT License)
- ✓ No dependencies required
- ✓ Works on all major platforms
- ✓ Active development and support

---

## Footer

### Footer Sections

#### Product
- Features
- Installation
- Documentation
- Changelog

#### Resources
- GitHub Repository
- Issue Tracker
- Contributing Guide
- Release Notes

#### Community
- Star on GitHub
- Report a Bug
- Request a Feature
- Contribute

#### Legal
- MIT License
- Privacy Policy (if applicable)
- Terms of Use (if applicable)

### Footer Credit
"Built with ❤️ by [MayR Labs](https://github.com/MayR-Labs)"

### Social Links
- GitHub: https://github.com/MayR-Labs
- Website: https://mayrlabs.com

---

## Technical Specifications

### Performance Requirements
- Page load time: < 3 seconds
- First contentful paint: < 1.5 seconds
- Mobile-responsive design (breakpoints: 640px, 768px, 1024px, 1280px)
- Accessibility: WCAG 2.1 Level AA compliance

### Browser Support
- Chrome (latest 2 versions)
- Firefox (latest 2 versions)
- Safari (latest 2 versions)
- Edge (latest 2 versions)

### Recommended Tech Stack
- **Static Site Generator:** Next.js, Astro, or Hugo
- **Styling:** Tailwind CSS
- **Hosting:** Vercel, Netlify, or GitHub Pages
- **Analytics:** Plausible or Google Analytics (privacy-friendly option preferred)

### Interactive Elements
- Terminal animation showing commands
- Copy-to-clipboard buttons for code examples
- Smooth scroll navigation
- Collapsible command reference sections

---

## SEO Metadata

### Page Title
"envdoc - Environment Variable CLI Tool | Manage, Validate & Transform .env Files"

### Meta Description
"envdoc is a powerful CLI tool for managing environment variables. Generate documentation, validate schemas, encrypt secrets, sync environments, and convert between formats. Free, open source, cross-platform."

### Keywords
- environment variables
- env file manager
- dotenv tool
- env file validator
- env file encryption
- CLI tool
- DevOps tools
- configuration management
- .env file
- environment configuration

### Open Graph Tags
```html
<meta property="og:title" content="envdoc - Tame Your Environment Variables">
<meta property="og:description" content="A powerful CLI tool for managing, validating, and transforming environment variable files.">
<meta property="og:type" content="website">
<meta property="og:url" content="https://envdoc.mayrlabs.com">
<meta property="og:image" content="https://envdoc.mayrlabs.com/og-image.png">
```

### Twitter Card
```html
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:title" content="envdoc - Tame Your Environment Variables">
<meta name="twitter:description" content="A powerful CLI tool for managing, validating, and transforming environment variable files.">
<meta name="twitter:image" content="https://envdoc.mayrlabs.com/twitter-card.png">
```

---

## Assets Needed

### Images
1. **Logo** - envdoc text logo (SVG, PNG)
2. **OG Image** - Open Graph image (1200x630px)
3. **Twitter Card** - Twitter sharing image (1200x628px)
4. **Favicon** - Multiple sizes (16x16, 32x32, 180x180, 192x192, 512x512)
5. **Feature Icons** - Icons for each feature (SVG preferred)
6. **Terminal Screenshot** - Terminal showing envdoc in action

### Animations (Optional)
1. **Terminal Typing Animation** - Animated demo of envdoc commands
2. **Command Output Animation** - Showing before/after results
3. **Feature Card Hover Effects** - Subtle interactions

### Downloads
1. **Installation Script** - Link to install.sh
2. **Binary Downloads** - Links to release assets
3. **Source Code** - Link to GitHub repository

---

## Implementation Checklist

### Phase 1: Foundation
- [ ] Set up project with chosen tech stack
- [ ] Create responsive layout structure
- [ ] Implement navigation component
- [ ] Add hero section with terminal visual

### Phase 2: Content
- [ ] Add features section with cards
- [ ] Implement use cases section
- [ ] Add installation section with code blocks
- [ ] Create command reference table

### Phase 3: Interactive Elements
- [ ] Add copy-to-clipboard functionality
- [ ] Implement terminal animation
- [ ] Add smooth scroll navigation
- [ ] Create collapsible sections

### Phase 4: Polish
- [ ] Add SEO metadata
- [ ] Implement Open Graph tags
- [ ] Add analytics
- [ ] Optimize for performance
- [ ] Test accessibility
- [ ] Cross-browser testing

### Phase 5: Launch
- [ ] Deploy to hosting platform
- [ ] Configure custom domain
- [ ] Set up SSL certificate
- [ ] Submit to search engines

---

## Appendix: Problem Statement Section

### The Problems We Solve

#### 1. The Onboarding Nightmare
A new developer joins your team. They clone the codebase, run the app, and... crash. They're missing 15 environment variables. After two hours of detective work, they finally have a working configuration.

**Solution:** `envdoc create-example` creates templates instantly.

#### 2. The Drift Problem
You have `.env`, `.env.staging`, `.env.production`. They started identical. Now one has 47 keys, another has 52. Nobody knows which keys are missing where.

**Solution:** `envdoc compare` and `envdoc sync` keep everything aligned.

#### 3. The Security Risk
You need to share production configuration. Do you email it? Post it in Slack? Each option makes your security team wince.

**Solution:** `envdoc encrypt` makes encryption trivial with AES-256.

#### 4. The Documentation Gap
What environment variables does your application need? If your answer is "let me check the code," you're not alone.

**Solution:** `envdoc create-schema` generates documentation automatically.

#### 5. The Format Chaos
Your application uses `.env` files, but your Kubernetes deployment needs YAML. Your Docker Compose wants JSON. You write custom scripts that break every few months.

**Solution:** `envdoc to` and `envdoc from` handle conversions seamlessly.

---

*This specifications document was created to help build an effective landing page for envdoc. For questions or suggestions, please open an issue in the repository.*

**Last Updated:** November 2025  
**Author:** MayR Labs
